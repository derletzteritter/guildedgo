package client

import (
	"bytes"
	"encoding/json"
	"log"
)

type event struct {
	Callback func(*Client, any)
	Type     *interface{}
}

func (c *Client) On(e string, callback func(client *Client, v any)) {
	c.events[e] = append(c.events[e], event{
		Callback: callback,
	})
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
	for _, event := range eventsCB {
		event.Callback(c, eventType)
	}
}
