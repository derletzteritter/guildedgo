package client

import (
	"bytes"
	"encoding/json"
	"log"
	"reflect"
)

type event struct {
	Callback func(*Client, any)
	Type     *interface{}
}

func (r *Client) On(e any, callback func(client *Client, v any)) {
	eventName := reflect.TypeOf(e).String()

	r.events[eventName] = append(r.events[eventName], event{
		Callback: callback,
	})
}

func (r *Client) onEvent(msg []byte) {
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
	eventsCB := r.events[re.T]
	for _, event := range eventsCB {
		event.Callback(r, eventType)
	}
}
