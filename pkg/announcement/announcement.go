package announcement

import (
	"fmt"
	"github.com/itschip/guildedgo/pkg/channel"
	"github.com/itschip/guildedgo/pkg/client"
	"net/http"
	"net/url"
	"strconv"
)

type Announcement struct {
	ID        string          `json:"id"`
	ServerID  string          `json:"serverId"`
	ChannelID string          `json:"channelId"`
	CreatedAt string          `json:"createdAt"`
	CreatedBy string          `json:"createdBy"`
	Content   string          `json:"content"`
	Mentions  channel.Mention `json:"mentions,omitempty"`
	Title     string          `json:"title"`
}

type CreateParams struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type GetParams struct {
	Before string
	Limit  int
}

func Create(c *client.Client, channelID string, params *CreateParams) (*Announcement, error) {
	var err error
	endpoint := client.GuildedApi + "/channels" + channelID + "/announcements"

	var v struct {
		Announcement `json:"announcement"`
	}

	res, err := c.Http.PerformRequest(http.MethodPost, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create announcement: %w", err)
	}

	err = c.Http.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Announcement, nil
}

func Get(c *client.Client, channelID string, params *GetParams) ([]Announcement, error) {
	var err error
	endpoint := client.GuildedApi + "/channels" + channelID + "/announcements"

	urlValues := url.Values{}

	if params.Before != "" {
		urlValues.Add("before", params.Before)
	}

	if params.Limit != 0 {
		urlValues.Add("limit", strconv.Itoa(params.Limit))
	}

	endpoint += "?" + urlValues.Encode()

	var v struct {
		Announcements []Announcement `json:"announcements"`
	}

	res, err := c.Http.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get announcements: %w", err)
	}

	err = c.Http.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return v.Announcements, nil
}

func Read(c *client.Client, channelID string, announcementID string) (*Announcement, error) {
	var err error
	endpoint := client.GuildedApi + "/channels" + channelID + "/announcements" + announcementID

	var v struct {
		Announcement `json:"announcement"`
	}

	res, err := c.Http.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get announcement: %w", err)
	}

	err = c.Http.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Announcement, nil
}

func Update(c *client.Client, channelID string, announcementID string, params *CreateParams) (*Announcement, error) {
	var err error
	endpoint := client.GuildedApi + "/channels" + channelID + "/announcements" + announcementID

	var v struct {
		Announcement `json:"announcement"`
	}

	res, err := c.Http.PerformRequest(http.MethodPatch, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update announcement: %w", err)
	}

	err = c.Http.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Announcement, nil
}

func Delete(c *client.Client, channelID string, announcementID string) error {
	var err error
	endpoint := client.GuildedApi + "/channels" + channelID + "/announcements" + announcementID

	_, err = c.Http.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to delete announcement: %w", err)
	}

	return nil
}
