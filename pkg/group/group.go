package group

import (
	"fmt"
	"io"
	"net/http"
)

type Client interface {
	PerformRequest(method, url string, data any) (io.ReadCloser, error)
	Decode(body io.ReadCloser, v any) error
}

const (
	guildedApi = "https://www.guilded.gg/api/v1"
)

// TODO: Add to socket interfaces
type Group struct {
	ID          string `json:"id"`
	ServerID    string `json:"serverId"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	IsHome      bool   `json:"isHome,omitempty"`
	EmoteID     int    `json:"emoteId,omitempty"`
	IsPublic    bool   `json:"isPublic,omitempty"`
	CreatedAt   string `json:"createdAt"`
	CreatedBy   string `json:"createdBy"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
	UpdatedBy   string `json:"updatedBy,omitempty"`
	ArchivedAt  string `json:"archivedAt,omitempty"`
	ArchivedBy  string `json:"archivedBy,omitempty"`
}

type groupResponse struct {
	Group `json:"group"`
}

type CreateParams struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	EmoteID     int    `json:"emoteId,omitempty"`
	IsPublic    bool   `json:"isPublic,omitempty"`
}

func Create(c Client, serverID string, params CreateParams) (Group, error) {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/groups"

	var v struct {
		Group `json:"group"`
	}

	body, err := c.PerformRequest(http.MethodPost, endpoint, params)
	if err != nil {
		return Group{}, fmt.Errorf("failed to create group: %w", err)
	}

	err = c.Decode(body, &v)
	if err != nil {
		return Group{}, fmt.Errorf("failed to decode group response: %w", err)
	}

	return v.Group, nil
}

// Get returns all groups in a server.
func Get(c Client, serverID string) ([]Group, error) {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/groups"

	var v struct {
		Groups []Group `json:"groups"`
	}

	body, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %w", err)
	}

	err = c.Decode(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode groups response: %w", err)
	}

	return v.Groups, nil
}

// Find returns a group in a server.
func Find(c Client, serverID, groupID string) (Group, error) {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/groups/" + groupID

	var v struct {
		Group `json:"group"`
	}

	body, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return Group{}, fmt.Errorf("failed to find group: %w", err)
	}

	err = c.Decode(body, &v)
	if err != nil {
		return Group{}, fmt.Errorf("failed to decode group response: %w", err)
	}

	return v.Group, nil
}

func Update(c Client, serverID, groupID string, params CreateParams) (Group, error) {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/groups/" + groupID

	var v struct {
		Group `json:"group"`
	}

	body, err := c.PerformRequest(http.MethodPatch, endpoint, params)
	if err != nil {
		return Group{}, fmt.Errorf("failed to update group: %w", err)
	}

	err = c.Decode(body, &v)
	if err != nil {
		return Group{}, fmt.Errorf("failed to decode group response: %w", err)
	}

	return v.Group, nil
}

func Delete(c Client, serverID, groupID string) error {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/groups/" + groupID

	_, err = c.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}

	return nil
}
