package channel

import (
	"fmt"
	"net/http"

	"github.com/itschip/guildedgo/pkg/client"
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
	// What users can access the channel.
	// Only applicable to server channels. If not present, this channel will respect normal permissions.
	// public is accessible to everyone, even those who aren't of the server.
	// private is only accessible to explicitly mentioned users.
	// Currently, threads cannot be public and other channels cannot be private.
	// Additionally, private threads can only exist with an associated messageId that is for a private message
	Visibility string `json:"visibility,omitempty"`
	ArchivedBy string `json:"archivedBy,omitempty"`
	ArchivedAt string `json:"archivedAt,omitempty"`
}

type Visibility string

const (
	VisibilityPublic  Visibility = "public"
	VisibilityPrivate Visibility = "private"
)

const (
	TypeAnnouncements string = "announcements"
	TypeChat          string = "chat"
	TypeCalendar      string = "calendar"
	TypeForums        string = "forums"
	TypeMedia         string = "media"
	TypeDocs          string = "docs"
	TypeVoice         string = "voice"
	TypeList          string = "list"
	TypeScheduling    string = "scheduling"
	TypeStream        string = "stream"
)

type Mention struct {
	// Info on mentioned users (min items 1)
	Users []MentionUser `json:"users,omitempty"`

	// Info on mentioned channels (min items 1)
	Channels []MentionChannel `json:"channels,omitempty"`

	// Info on mentioned roles (min items 1)
	Roles []MentionRole `json:"roles,omitempty"`

	// If @everyone was mentioned
	Everyone bool `json:"everyone,omitempty"`

	// If @here was mentioned
	Here bool `json:"here,omitempty"`
}

type MentionRole struct {
	// The ID of the role
	ID int `json:"id"`
}

type MentionChannel struct {
	// The ID of the channel
	ID string `json:"id"`
}

type MentionUser struct {
	// The ID of the user
	ID string `json:"id"`
}

type CreateParams struct {
	Name  string `json:"name"`
	Topic string `json:"topic,omitempty"`
	// "private" or "public". optional
	Visibility Visibility `json:"visibility,omitempty"`
	Type       string     `json:"type"`
	ServerID   string     `json:"serverId,omitempty"`
	GroupID    string     `json:"groupId,omitempty"`
	CategoryID int        `json:"categoryId,omitempty"`
	ParentID   string     `json:"parentId,omitempty"`
	MessageID  string     `json:"messageId,omitempty"`
}

type UpdateParams struct {
	Name       string     `json:"name,omitempty"`
	Topic      string     `json:"topic,omitempty"`
	Visibility Visibility `json:"visibility,omitempty"`
}

func Create(c *client.Client, params *CreateParams) (ServerChannel, error) {
	var err error
	endpoint := client.GuildedApi + "/channels"

	var v struct {
		Channel ServerChannel `json:"channel"`
	}

	body, err := c.Http.PerformRequest(http.MethodPost, endpoint, params)
	if err != nil {
		return ServerChannel{}, fmt.Errorf("failed to create channel: %w", err)
	}

	err = c.Http.Decode(body, &v)
	if err != nil {
		return ServerChannel{}, fmt.Errorf("failed to decode channel response: %w", err)
	}

	return v.Channel, nil
}

func Get(c *client.Client, channelID string) (ServerChannel, error) {
	var err error
	endpoint := client.GuildedApi + "/channels/" + channelID

	var v struct {
		Channel ServerChannel `json:"channel"`
	}

	body, err := c.Http.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return ServerChannel{}, fmt.Errorf("failed to get channel: %w", err)
	}

	err = c.Http.Decode(body, &v)
	if err != nil {
		return ServerChannel{}, fmt.Errorf("failed to decode channel response: %w", err)
	}

	return v.Channel, nil
}

func Update(c *client.Client, channelID string, params *UpdateParams) (ServerChannel, error) {
	var err error
	endpoint := client.GuildedApi + "/channels/" + channelID

	var v struct {
		Channel ServerChannel `json:"channel"`
	}

	body, err := c.Http.PerformRequest(http.MethodPatch, endpoint, params)
	if err != nil {
		return ServerChannel{}, fmt.Errorf("failed to update channel: %w", err)
	}

	err = c.Http.Decode(body, &v)
	if err != nil {
		return ServerChannel{}, fmt.Errorf("failed to decode channel response: %w", err)
	}

	return v.Channel, nil
}

func Delete(c *client.Client, channelID string) error {
	var err error
	endpoint := client.GuildedApi + "/channels/" + channelID

	_, err = c.Http.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to delete channel: %w", err)
	}

	return nil
}

func Archive(c *client.Client, channelID string) error {
	var err error
	endpoint := client.GuildedApi + "/channels/" + channelID + "/archive"

	_, err = c.Http.PerformRequest(http.MethodPut, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to archive channel: %w", err)
	}

	return nil
}

func Restore(c *client.Client, channelID string) error {
	var err error
	endpoint := client.GuildedApi + "/channels/" + channelID + "/archive"

	_, err = c.Http.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to restore channel: %w", err)
	}

	return nil
}
