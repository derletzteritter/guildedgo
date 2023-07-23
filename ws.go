package guildedgo

import (
	"bytes"
	"encoding/json"
	"log"
)

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
	interfaces["ChannelArchived"] = &ChannelArchived{}
	interfaces["ChannelRestored"] = &ChannelRestored{}
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
