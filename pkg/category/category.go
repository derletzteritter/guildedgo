package category

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

type Category struct {
	ID        int    `json:"id"`
	ServerID  string `json:"serverId"`
	GroupID   string `json:"groupId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	Name      string `json:"name"`
}

type CreateParams struct {
	Name    string `json:"name"`
	GroupID string `json:"groupId,omitempty"`
}

type UpdateParams struct {
	Name string `json:"name"`
}

func Create(c Client, serverID string, params *CreateParams) (Category, error) {
	endpoint := guildedApi + "/servers/" + serverID + "/categories"

	var v struct {
		Category `json:"category"`
	}

	body, err := c.PerformRequest(http.MethodPost, endpoint, params)
	if err != nil {
		return Category{}, fmt.Errorf("failed to create category: %w", err)
	}

	err = c.Decode(body, &v)
	if err != nil {
		return Category{}, fmt.Errorf("failed to decode category response: %w", err)
	}

	return v.Category, nil
}

func Read(c Client, serverID string, categoryID string) (Category, error) {
	endpoint := guildedApi + "/servers" + serverID + "/categories/" + categoryID

	var v struct {
		Category `json:"category"`
	}

	body, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return Category{}, fmt.Errorf("failed to get category: %w", err)
	}

	err = c.Decode(body, &v)
	if err != nil {
		return Category{}, fmt.Errorf("failed to decode category response: %w", err)
	}

	return v.Category, nil
}

func Update(c Client, serverID string, categoryID int, params *UpdateParams) (Category, error) {
	endpoint := fmt.Sprintf("%s/servers/%s/categories/%d", guildedApi, serverID, categoryID)

	var v struct {
		Category `json:"category"`
	}

	body, err := c.PerformRequest(http.MethodPatch, endpoint, params)
	if err != nil {
		return Category{}, fmt.Errorf("failed to update category: %w", err)
	}

	err = c.Decode(body, &v)
	if err != nil {
		return Category{}, fmt.Errorf("failed to decode category response: %w", err)
	}

	return v.Category, nil
}

func Delete(c Client, serverID string, categoryID int) (Category, error) {
	endpoint := fmt.Sprintf("%s/servers/%s/categories/%d", guildedApi, serverID, categoryID)

	var v struct {
		Category `json:"category"`
	}

	body, err := c.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return Category{}, fmt.Errorf("failed to delete category: %w", err)
	}

	err = c.Decode(body, &v)
	if err != nil {
		return Category{}, fmt.Errorf("failed to decode category response: %w", err)
	}

	return v.Category, nil
}
