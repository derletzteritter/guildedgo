package category

import (
	"fmt"
	"net/http"

	"github.com/itschip/guildedgo/pkg/client"
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

func Create(c *client.Client, params *CreateParams) (Category, error) {
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/categories"

	var v struct {
		Category `json:"category"`
	}

	body, err := c.Http.PerformRequest(http.MethodPost, endpoint, params)
	if err != nil {
		return Category{}, fmt.Errorf("failed to create category: %w", err)
	}

	err = c.Http.Decode(body, &v)
	if err != nil {
		return Category{}, fmt.Errorf("failed to decode category response: %w", err)
	}

	return v.Category, nil
}

func Read(c *client.Client, categoryID string) (Category, error) {
	endpoint := client.GuildedApi + "/servers" + c.ServerID + "/categories/" + categoryID

	var v struct {
		Category `json:"category"`
	}

	body, err := c.Http.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return Category{}, fmt.Errorf("failed to get category: %w", err)
	}

	err = c.Http.Decode(body, &v)
	if err != nil {
		return Category{}, fmt.Errorf("failed to decode category response: %w", err)
	}

	return v.Category, nil
}

func Update(c *client.Client, categoryID int, params *UpdateParams) (Category, error) {
	endpoint := fmt.Sprintf("%s/servers/%s/categories/%d", client.GuildedApi, c.ServerID, categoryID)

	var v struct {
		Category `json:"category"`
	}

	body, err := c.Http.PerformRequest(http.MethodPatch, endpoint, params)
	if err != nil {
		return Category{}, fmt.Errorf("failed to update category: %w", err)
	}

	err = c.Http.Decode(body, &v)
	if err != nil {
		return Category{}, fmt.Errorf("failed to decode category response: %w", err)
	}

	return v.Category, nil
}

func Delete(c *client.Client, categoryID int) (Category, error) {
	endpoint := fmt.Sprintf("%s/servers/%s/categories/%d", client.GuildedApi, c.ServerID, categoryID)

	var v struct {
		Category `json:"category"`
	}

	body, err := c.Http.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return Category{}, fmt.Errorf("failed to delete category: %w", err)
	}

	err = c.Http.Decode(body, &v)
	if err != nil {
		return Category{}, fmt.Errorf("failed to decode category response: %w", err)
	}

	return v.Category, nil
}
