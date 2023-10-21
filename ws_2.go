package guildedgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type RawEvent2 struct {
	OP   int             `json:"op"`
	T    string          `json:"t"`
	S    string          `json:"s"`
	Data json.RawMessage `json:"d"`
}

type WelcomeOP struct {
	HeartbeatInterval int `json:"heartbeatIntervalMs"`
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (c *Client) Open() {
	var err error

	c.Lock()
	defer c.Unlock()

	if c.conn != nil {
		return
	}

	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	c.conn, _, err = websocket.DefaultDialer.Dial("wss://www.guilded.gg/websocket/v1", header)
	if err != nil {
		log.Fatalln("Failed to connect to websocket: ", err.Error())
	}

	c.conn.SetCloseHandler(func(code int, text string) error {
		return nil
	})

	defer func() {
		if err != nil {
			c.conn.Close()
			c.conn = nil
		}
	}()

	// Get the welcome message from Guilded
	_, m, err := c.conn.ReadMessage()
	if err != nil {
		log.Fatalln("Failed to read message: ", err.Error())
	}
	m = bytes.TrimSpace(bytes.Replace(m, newline, space, -1))

	event := c.onWelcomeMessage(m)
	if event == nil {
		return
	}

	c.listening = make(chan struct{})

	go c.heartbeat(c.conn, c.listening, event.HeartbeatInterval)
	go c.listen(c.conn, c.listening)

	log.Println("Listening for messages in main")
}

func (c *Client) listen(wsConn *websocket.Conn, listening <-chan struct{}) {
	for {
		select {
		case <-listening:
			return
		default:
			_, msg, err := wsConn.ReadMessage()
			if err != nil {
				log.Println("Failed to read message: ", err.Error())
				return
			}

			msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))

			c.onEvent(msg)
		}
	}
}

func (c *Client) heartbeat(wsConn *websocket.Conn, listening <-chan struct{}, intervalMs int) {
	ticker := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Sending heartbeat")
			err := wsConn.WriteMessage(websocket.TextMessage, []byte(""))
			if err != nil {
				log.Println("Failed to send heartbeat: ", err.Error())
			}
		case <-listening:
			return
		}
	}
}

func (c *Client) Close() {
	c.Lock()

	if c.listening != nil {
		log.Println("Closing listening channel")
		close(c.listening)
		c.listening = nil
	}

	if c.conn != nil {
		log.Println("Closing websocket connection")

		c.wsMutex.Lock()
		err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		c.wsMutex.Unlock()
		if err != nil {
			log.Println("Failed to write close message: ", err.Error())
		}

		time.Sleep(1 * time.Second)

		log.Println("Closing websocket connection")

		err = c.conn.Close()
		if err != nil {
			log.Println("Failed to close websocket connection: ", err.Error())
		}

		c.conn = nil
	}

	c.Unlock()

	log.Println("Closed websocket connection")

	return
}

func (c *Client) onWelcomeMessage(msg []byte) *WelcomeOP {
	var err error
	reader := bytes.NewBuffer(msg)

	var re RawEvent2
	decoder := json.NewDecoder(reader)

	err = decoder.Decode(&re)
	if err != nil {
		log.Println("Failed to decode raw event")
	}

	if re.OP != 1 {
		log.Println("Expected OP code 1, got", re.OP)
		return nil
	}

	var h WelcomeOP
	err = json.Unmarshal(re.Data, &h)
	if err != nil {
		log.Printf("Failed to unmarshal event data for %q. Error: %s", re.T, err.Error())
	}

	return &h
}
