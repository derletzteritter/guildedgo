package guildedgo

import (
	"net/url"
	"strconv"
)

type Announcement struct {
	ID        string   `json:"id"`
	ServerID  string   `json:"serverId"`
	ChannelID string   `json:"channelId"`
	CreatedAt string   `json:"createdAt"`
	CreatedBy string   `json:"createdBy"`
	Content   string   `json:"content"`
	Mentions  Mentions `json:"mentions,omitempty"`
	Title     string   `json:"title"`
}

type GetAnnouncementParams struct {
	Before string
	Limit  int
}

type AnncounmentResponse struct {
	Announcement `json:"announcement"`
}

type AnnouncementService interface {
	CreateAnnouncement(serverID string, channelID string, title string, content string) (*Announcement, error)
	GetAnnouncements(channelID string, query *GetAnnouncementParams) ([]Announcement, error)
	ReadAnnouncement(channelID string, announcementID string) (*Announcement, error)
	UpdateAnnouncement(channelID string, announcementID string, title string, content string) (*Announcement, error)
	DeleteAnnouncement(channelID string, announcementID string) error
}

type announcementEndpoints struct{}

func (endpoints *announcementEndpoints) Get(channelID string) string {
	return guildedApi + "/channels/" + channelID + "/announcements"
}

type announcementService struct {
	client    *Client
	endpoints *announcementEndpoints
}

var _ AnnouncementService = &announcementService{}

func (service *announcementService) CreateAnnouncement(serverID string, channelID string, title string, content string) (*Announcement, error) {
	endpoint := service.endpoints.Get(channelID)

	body := map[string]interface{}{
		"title":   title,
		"content": content,
	}

	var response AnncounmentResponse
	err := service.client.PostRequestV2(endpoint, body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Announcement, nil
}

func (service *announcementService) GetAnnouncements(channelID string, query *GetAnnouncementParams) ([]Announcement, error) {
	endpoint := service.endpoints.Get(channelID)

	params := url.Values{}

	if query != nil {
		if query.Before != "" {
			params.Add("before", query.Before)
		}

		if query.Limit != 0 {
			params.Add("limit", strconv.Itoa(query.Limit))
		}
	}

	endpoint = endpoint + "?" + params.Encode()

	var response struct {
		Announcements []Announcement `json:"announcements"`
	}

	err := service.client.GetRequestV2(endpoint, &response)
	if err != nil {
		return nil, err
	}

	return response.Announcements, nil
}

func (service *announcementService) ReadAnnouncement(channelID string, announcementID string) (*Announcement, error) {
	endpoint := service.endpoints.Get(channelID) + "/" + announcementID

	var response AnncounmentResponse
	err := service.client.GetRequestV2(endpoint, &response)
	if err != nil {
		return nil, err
	}

	return &response.Announcement, nil
}

func (service *announcementService) UpdateAnnouncement(channelID string, announcementID string, title string, content string) (*Announcement, error) {
	endpoint := service.endpoints.Get(channelID) + "/" + announcementID

	body := map[string]interface{}{
		"title":   title,
		"content": content,
	}

	var response AnncounmentResponse
	err := service.client.PatchRequest(endpoint, body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Announcement, nil
}

func (service *announcementService) DeleteAnnouncement(channelID string, announcementID string) error {
	endpoint := service.endpoints.Get(channelID) + "/" + announcementID

	_, err := service.client.DeleteRequest(endpoint)
	if err != nil {
		return err
	}

	return nil
}
