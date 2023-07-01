package guildedgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

type ServerChannel struct {
	ID string `json:"id"`

	// The type of channel. This will determine what routes to use for creating content in a channel.
	// For example, if this "chat", then one must use the routes for creating channel messages
	Type string `json:"type"`

	Name string `json:"name"`

	// The topic of the channel. Not applicable to threads (min length 1; max length 512)
	Topic string `json:"topic,omitempty"`

	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	ServerID  string `json:"serverId"`

	// ID of the root channel or thread in the channel hierarchy.
	// Only applicable to "chat", "voice", and "stream" channels and indicates that this channel is a thread, if present
	RootID string `json:"rootId,omitempty"`

	// ID of the immediate parent channel or thread in the channel hierarchy.
	// Only applicable to "chat", "voice", and "stream" channels and indicates that this channel is a thread, if present
	ParentID string `json:"parentId,omitempty"`

	// The ID of the message that this channel was created off of.
	// Only applicable to "chat", "voice", and "stream" channels and indicates that this channel is a thread, if present
	MessageID string `json:"messageId,omitempty"`

	// The category that the channel exists in. Only relevant for server channels
	CategoryID int    `json:"categoryId,omitempty"`
	GroupID    string `json:"groupId"`
	IsPublic   bool   `json:"isPublic,omitempty"`
	ArchivedBy string `json:"archivedBy,omitempty"`
	ArchivedAt string `json:"archivedAt,omitempty"`
}

const (
	ChannelTypeAnnouncements string = "announcements"
	ChannelTypeChat          string = "chat"
	ChannelTypeCalendar      string = "calendar"
	ChannelTypeForums        string = "forums"
	ChannelTypeMedia         string = "media"
	ChannelTypeDocs          string = "docs"
	ChannelTypeVoice         string = "voice"
	ChannelTypeList          string = "list"
	ChannelTypeScheduling    string = "scheduling"
	ChannelTypeStream        string = "stream"
)

type Mentions struct {
	// Info on mentioned users (min items 1)
	Users []MentionsUser `json:"users,omitempty"`

	// Info on mentioned channels (min items 1)
	Channels []MentionsChannel `json:"channels,omitempty"`

	// Info on mentioned roles (min items 1)
	Roles []MentionsRole `json:"roles,omitempty"`

	// If @everyone was mentioned
	Everyone bool `json:"everyone,omitempty"`

	// If @here was mentioned
	Here bool `json:"here,omitempty"`
}

type MentionsRole struct {
	// The ID of the role
	ID int `json:"id"`
}

type MentionsChannel struct {
	// The ID of the channel
	ID string `json:"id"`
}

type MentionsUser struct {
	// The ID of the user
	ID string `json:"id"`
}

type NewChannelObject struct {
	Name       string `json:"name"`
	Topic      string `json:"topic,omitempty"`
	IsPublic   bool   `json:"isPublic,omitempty"`
	Type       string `json:"type"`
	ServerID   string `json:"serverId,omitempty"`
	GroupID    string `json:"groupId,omitempty"`
	CategoryID int    `json:"categoryId,omitempty"`
	ParentID   string `json:"parentId,omitempty"`
	MessageID  string `json:"messageId,omitempty"`
}

type UpdateChannelObject struct {
	Name     string `json:"name,omitempty"`
	Topic    string `json:"topic,omitempty"`
	IsPublic bool   `json:"isPublic,omitempty"`
}

type ServerChannelResponse struct {
	Channel ServerChannel `json:"channel"`
}

type ChannelService interface {
	CreateChannel(channelObject *NewChannelObject) (*ServerChannel, error)
	GetChannel(channelId string) (*ServerChannel, error)
	UpdateChannel(channelId string, channelObject *UpdateChannelObject) (*ServerChannel, error)
	DeleteChannel(channelId string) error
	SendMessage(channelId string, message *MessageObject) (*ChatMessage, error)
	GetMessages(channelId string, getObject *GetMessagesObject) (*[]ChatMessage, error)
	GetMessage(channelId string, messageId string) (*ChatMessage, error)
	UpdateChannelMessage(channelId string, messageId string, newMessage *MessageObject) (*ChatMessage, error)
	DeleteChannelMessage(channelId string, messageId string) error
}

type channelEndpoints struct{}

func (e *channelEndpoints) Default() string {
	return guildedApi + "/channels"
}

func (e *channelEndpoints) Get(channelId string) string {
	return guildedApi + "/channels/" + channelId
}

func (e *channelEndpoints) Message(channelId string) string {
	return guildedApi + "/channels/" + channelId + "/messages"
}

func (e *channelEndpoints) ChannelMessages(channelId string) string {
	return guildedApi + "/channels/" + channelId + "/messages"
}

func (e *channelEndpoints) MessageID(channelId string, messageId string) string {
	return guildedApi + "/channels/" + channelId + "/messages/" + messageId
}

type channelService struct {
	client    *Client
	endpoints *channelEndpoints
}

var _ ChannelService = &channelService{}

// CreateChannel returns the newly created channel.
// Only server channels are supported at this time (coming soon™: DM Channels!)
func (service *channelService) CreateChannel(channelObject *NewChannelObject) (*ServerChannel, error) {
	endpoint := service.endpoints.Default()

	channelObject.ServerID = service.client.ServerID

	resp, err := service.client.PostRequest(endpoint, &channelObject)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create new channel. Error: \n%v", err.Error()))
	}

	var serverChannel ServerChannelCreated
	err = json.Unmarshal(resp, &serverChannel)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to unmarshal ServerChannel response. Error: \n%v", err.Error()))
	}

	return &serverChannel.Channel, nil
}

// GetChannel returns a channel by channelId.
// Only server channels are supported at this time (coming soon™: DM Channels!)
func (service *channelService) GetChannel(channelId string) (*ServerChannel, error) {
	endpoint := service.endpoints.Get(channelId)

	var serverChannel ServerChannelResponse
	err := service.client.GetRequestV2(endpoint, &serverChannel)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get channel. Error: \n%v", err.Error()))
	}

	return &serverChannel.Channel, nil
}

// UpdateChannel returns the updated channel.
func (service *channelService) UpdateChannel(channelId string, channelObject *UpdateChannelObject) (*ServerChannel, error) {
	endpoint := service.endpoints.Get(channelId)

	var serverChannel ServerChannelResponse
	err := service.client.PatchRequest(endpoint, &channelObject, &serverChannel)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to update  channel. Error: \n%v", err.Error()))
	}

	return &serverChannel.Channel, nil
}

// DeleteChannel does not return anything
// Only server channels are supported at this time (coming soon™: DM Channels!)
func (cs *channelService) DeleteChannel(channelId string) error {
	endpoint := cs.endpoints.Get(channelId)
	_, err := cs.client.DeleteRequest(endpoint)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to delete channel. Error: \n%v", err.Error()))
	}

	return nil
}

func (service *channelService) SendMessage(channelId string, message *MessageObject) (*ChatMessage, error) {
	endpoint := service.endpoints.Message(channelId)

	resp, err := service.client.PostRequest(endpoint, &message)
	if err != nil {
		return nil, err
	}

	var msg MessageResponse
	err = json.Unmarshal(resp, &msg)
	if err != nil {
		return nil, err
	}

	return &msg.Message, err
}

// TODO: only allow for content and embed updates

func (service *channelService) UpdateChannelMessage(channelId string, messageId string, newMessage *MessageObject) (*ChatMessage, error) {
	endpoint := service.endpoints.MessageID(channelId, messageId)

	resp, err := service.client.PutRequest(endpoint, &newMessage)
	if err != nil {
		return nil, err
	}

	var msg MessageResponse
	err = json.Unmarshal(resp, &msg)
	if err != nil {
		return nil, err
	}

	return &msg.Message, err
}

// GetMessages TODO: add support for params
func (service *channelService) GetMessages(channelId string, getObject *GetMessagesObject) (*[]ChatMessage, error) {
	endpoint := service.endpoints.ChannelMessages(channelId)

	// create query params with getObject
	params := url.Values{}

	if getObject != nil {
		if getObject.Before != "" {
			params.Add("before", getObject.Before)
		}
		if getObject.After != "" {
			params.Add("after", getObject.After)
		}

		if getObject.Limit != 0 {
			params.Add("limit", strconv.Itoa(getObject.Limit))
		}

		if getObject.IncludePrivate {
			params.Add("includePrivate", strconv.FormatBool(getObject.IncludePrivate))
		}
	}

	endpoint = endpoint + "?" + params.Encode()

	var msgs AllMessagesResponse
	err := service.client.GetRequestV2(endpoint, &msgs)
	if err != nil {
		return nil, err
	}

	return &msgs.Messages, nil
}

// GetMessage Get a message from a channel
func (service *channelService) GetMessage(channelId string, messageId string) (*ChatMessage, error) {
	endpoint := service.endpoints.MessageID(channelId, messageId)

	resp, err := service.client.GetRequest(endpoint)
	if err != nil {
		return nil, err
	}

	var msg MessageResponse
	err = json.Unmarshal(resp, &msg)
	if err != nil {
		return nil, err
	}

	return &msg.Message, nil
}

func (service *channelService) DeleteChannelMessage(channelId string, messageId string) error {
	endpoint := service.endpoints.MessageID(channelId, messageId)

	_, err := service.client.DeleteRequest(endpoint)
	if err != nil {
		return err
	}

	return nil
}
