package channel

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Announcement struct {
	ID        string  `json:"id"`
	ServerID  string  `json:"serverId"`
	ChannelID string  `json:"channelId"`
	CreatedAt string  `json:"createdAt"`
	CreatedBy string  `json:"createdBy"`
	Content   string  `json:"content"`
	Mentions  Mention `json:"mentions,omitempty"`
	Title     string  `json:"title"`
}

type Comment struct {
	ID             int     `json:"id"`
	Content        string  `json:"content"`
	CreatedAt      string  `json:"createdAt"`
	UpdatedAt      string  `json:"updatedAt"`
	CreatedBy      string  `json:"createdBy"`
	ChannelID      string  `json:"channelId"`
	AnnouncementID string  `json:"announcementId"`
	Mentions       Mention `json:"mentions,omitempty"`
}

type CreateAnnouncementParams struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type GetAnnouncementParams struct {
	Before string
	Limit  int
}

func CreateAnnouncement(c Client, channelID string, params *CreateAnnouncementParams) (*Announcement, error) {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/announcements"

	var v struct {
		Announcement `json:"announcement"`
	}

	res, err := c.PerformRequest(http.MethodPost, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create announcement: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Announcement, nil
}

func GetAnnouncement(c Client, channelID string, params *GetAnnouncementParams) ([]Announcement, error) {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/announcements"

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

	res, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get announcements: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return v.Announcements, nil
}

func ReadAnnouncement(c Client, channelID string, announcementID string) (*Announcement, error) {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/announcements" + announcementID

	var v struct {
		Announcement `json:"announcement"`
	}

	res, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get announcement: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Announcement, nil
}

func UpdateAnnouncement(c Client, channelID string, announcementID string, params *CreateAnnouncementParams) (*Announcement, error) {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/announcements" + announcementID

	var v struct {
		Announcement `json:"announcement"`
	}

	res, err := c.PerformRequest(http.MethodPatch, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update announcement: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Announcement, nil
}

func DeleteAnnouncement(c Client, channelID string, announcementID string) error {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/announcements" + announcementID

	_, err = c.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to delete announcement: %w", err)
	}

	return nil
}

func CreateComment(c Client, channelID string, announcementID string, params *CreateAnnouncementParams) (*Comment, error) {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/announcements" + announcementID + "/comments"

	var v struct {
		Comment `json:"announcementComment"`
	}

	res, err := c.PerformRequest(http.MethodPost, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create comment: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Comment, nil
}

func GetComments(c Client, channelID string, announcementID string) ([]Comment, error) {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/announcements" + announcementID + "/comments"

	var v struct {
		Comments []Comment `json:"announcementComments"`
	}

	res, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return v.Comments, nil
}

func FindComment(c Client, channelID string, announcementID string, commentID int) (*Comment, error) {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/announcements" + announcementID + "/comments" + strconv.Itoa(commentID)

	var v struct {
		Comment `json:"announcementComment"`
	}

	res, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get comment: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Comment, nil
}

func UpdateComment(c Client, channelID string, announcementID string, commentID int, params *CreateAnnouncementParams) (*Comment, error) {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/announcements" + announcementID + "/comments" + strconv.Itoa(commentID)

	var v struct {
		Comment `json:"announcementComment"`
	}

	res, err := c.PerformRequest(http.MethodPatch, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update comment: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Comment, nil
}

func DeleteComment(c Client, channelID string, announcementID string, commentID int) error {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/announcements" + announcementID + "/comments" + strconv.Itoa(commentID)

	_, err = c.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}

	return nil
}
