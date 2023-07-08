package guildedgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (c *Client) Open() {
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	conn, _, err := websocket.DefaultDialer.Dial("wss://www.guilded.gg/websocket/v1", header)
	c.conn = conn

	if err != nil {
		log.Fatalln("Failed to connect to websocket: ", err.Error())
	}

	defer func() {
		err := c.conn.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()

	_, m, err := c.conn.ReadMessage()
	if err != nil {
		log.Fatalln("Failed to read message: ", err.Error())
	}

	m = bytes.TrimSpace(bytes.Replace(m, newline, space, -1))

	listening := make(chan struct{})

	go func() {
		defer close(listening)

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err.Error())
				return
			}

			c.onEvent(msg)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-listening:
			return

		case t := <-ticker.C:
			err := c.conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				return
			}

		case <-interrupt:
			log.Println("Interrupt")

			err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				return
			}

			select {
			case <-listening:
			case <-time.After(time.Second):
			}

			return
		}
	}
}

func (c *Client) Close() {
	c.events = make(map[string][]Event) // Clear registered event callbacks

	// Close the WebSocket connection
	if c.conn != nil {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("Failed to send close message:", err.Error())
		}

		select {
		case <-interrupt:
			// Received interrupt signal, handle it gracefully
			log.Println("Interrupt signal received while closing the connection.")

			err := c.conn.Close()
			if err != nil {
				log.Println("Failed to close WebSocket connection:", err.Error())
			}
			case <-time.After(time.Second):
				// Timeout for graceful termination reached, force close the connection
				log.Println("Timeout reached for graceful termination. Forcing connection close.")

				err := c.conn.Close()
				if err != nil {
					log.Println("Failed to close WebSocket connection:", err.Error())
				}
		}

		c.conn = nil
	}
}

var interfaces = make(map[string]any)

func init() {
	interfaces["BotServerMembershipCreated"] = &BotServerMembershipCreated{}
	interfaces["BotServerMembershipDeleted"] = &BotServerMembershipDeleted{}
	interfaces["ChatMessageCreated"] = &ChatMessageCreated{}
	interfaces["ChatMessageUpdated"] = &ChatMessageUpdated{}
	interfaces["ChatMessageDeleted"] = &ChatMessageDeleted{}
	interfaces["ServerMemberJoined"] = &ServerMemberJoined{}
	interfaces["ServerMemberRemoved"] = &ServerMemberRemoved{}
	interfaces["ServerMemberBanned"] = &ServerMemberBanned{}
	interfaces["ServerMemberUnbanned"] = &ServerMemberUnbanned{}
	interfaces["ServerMemberUpdated"] = &ServerMemberUpdated{}
	interfaces["ServerRolesUpdated"] = &ServerRolesUpdated{}
	interfaces["ServerChannelCreated"] = &ServerChannelCreated{}
	interfaces["ServerChannelUpdated"] = &ServerChannelUpdated{}
	interfaces["ServerChannelDeleted"] = &ServerChannelDeleted{}
	interfaces["ServerWebhookCreated"] = &ServerWebhookCreated{}
	interfaces["ServerWebhookUpdated"] = &ServerWebhookUpdated{}
	interfaces["DocCreated"] = &DocCreated{}
	interfaces["DocUpdated"] = &DocUpdated{}
	interfaces["DocDeleted"] = &DocDeleted{}
	interfaces["CalendarEventCreated"] = &CalendarEventCreated{}
	interfaces["CalendarEventUpdated"] = &CalendarEventUpdated{}
	interfaces["CalendarEventDeleted"] = &CalendarEventDeleted{}
	interfaces["ForumTopicCreated"] = &ForumTopicCreated{}
	interfaces["ForumTopicUpdated"] = &ForumTopicUpdated{}
	interfaces["ForumTopicDeleted"] = &ForumTopicDeleted{}
	interfaces["ForumTopicPinned"] = &ForumTopicPinned{}
	interfaces["ForumTopicUnpinned"] = &ForumTopicUnpinned{}
	interfaces["ForumTopicReactionCreated"] = &ForumTopicReactionCreated{}
	interfaces["ForumTopicReactionDeleted"] = &ForumTopicReactionDeleted{}
	interfaces["ForumTopicLocked"] = &ForumTopic{}
	interfaces["ForumTopicUnlocked"] = &ForumTopic{}
	interfaces["ForumTopicCommentCreated"] = &ForumTopicComment{}
	interfaces["ForumTopicCommentUpdated"] = &ForumTopicComment{}
	interfaces["ForumTopicCommentDeleted"] = &ForumTopicComment{}
	interfaces["CalendarEventRsvpUpdated"] = &CalendarEventRsvp{}
	interfaces["CalendarEventRsvpManyUpdated"] = &[]CalendarEventRsvp{}
	interfaces["CalendarEventRsvpDeleted"] = &CalendarEventRsvp{}
}

type RawEvent struct {
	T    string          `json:"t"`
	S    string          `json:"s"`
	Data json.RawMessage `json:"d"`
}

func (c *Client) onEvent(msg []byte) {
	var err error
	reader := bytes.NewBuffer(msg)

	var re *RawEvent
	decoder := json.NewDecoder(reader)

	err = decoder.Decode(&re)
	if err != nil {
		log.Println("Failed to decode raw event")
	}

	eventType := interfaces[re.T]
	err = json.Unmarshal(re.Data, eventType)
	if err != nil {
		log.Printf("Failed to unmarshal event data for %q. Error: %s", re.T, err.Error())
	}

	// Is this smart? Probably not.
	eventsCB := c.events[re.T]
	for _, callback := range eventsCB {
		callback.Callback(c, eventType)
	}
}
