package group

import (
	"fmt"
	"net/http"

	"github.com/itschip/guildedgo/pkg/client"
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

func Create(c *client.Client, params CreateParams) (Group, error) {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/groups"

	var v struct {
		Group `json:"group"`
	}

	body, err := c.Http.PerformRequest(http.MethodPost, endpoint, params)
	if err != nil {
		return Group{}, fmt.Errorf("failed to create group: %w", err)
	}

	err = c.Http.Decode(body, &v)
	if err != nil {
		return Group{}, fmt.Errorf("failed to decode group response: %w", err)
	}

	return v.Group, nil
}

// Get returns all groups in a server.
func Get(c *client.Client) ([]Group, error) {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/groups"

	var v struct {
		Groups []Group `json:"groups"`
	}

	body, err := c.Http.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %w", err)
	}

	err = c.Http.Decode(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode groups response: %w", err)
	}

	return v.Groups, nil
}

// Find returns a group in a server.
func Find(c *client.Client, groupID string) (Group, error) {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/groups/" + groupID

	var v struct {
		Group `json:"group"`
	}

	body, err := c.Http.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return Group{}, fmt.Errorf("failed to find group: %w", err)
	}

	err = c.Http.Decode(body, &v)
	if err != nil {
		return Group{}, fmt.Errorf("failed to decode group response: %w", err)
	}

	return v.Group, nil
}

func Update(c *client.Client, groupID string, params CreateParams) (Group, error) {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/groups/" + groupID

	var v struct {
		Group `json:"group"`
	}

	body, err := c.Http.PerformRequest(http.MethodPatch, endpoint, params)
	if err != nil {
		return Group{}, fmt.Errorf("failed to update group: %w", err)
	}

	err = c.Http.Decode(body, &v)
	if err != nil {
		return Group{}, fmt.Errorf("failed to decode group response: %w", err)
	}

	return v.Group, nil
}

func Delete(c *client.Client, groupID string) error {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/groups/" + groupID

	_, err = c.Http.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}

	return nil
}
