package guildedgo

import (
	"fmt"
)

type Category struct {
	ID        int    `json:"id"`
	ServerID  string `json:"serverId"`
	GroupID   string `json:"groupId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Name      string `json:"name"`
}

type CreateCategory struct {
	Name    string `json:"name"`
	GroupID string `json:"groupId,omitempty"`
}

type CategoryService interface {
	Read(categoryID int) (*Category, error)
	Create(options *CreateCategory) (*Category, error)
	Update(categoryID int, name string) (*Category, error)
	Delete(categoryID int) error
}

type categoryService struct {
	client *Client
}

var _ CategoryService = &categoryService{}

func (s *categoryService) Read(categoryID int) (*Category, error) {
	endpoint := fmt.Sprintf("%s/servers/%s/categories/%d", guildedApi, s.client.ServerID, categoryID)

	var category Category
	err := s.client.GetRequestV2(endpoint, &category)
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	return &category, nil
}

func (s *categoryService) Create(options *CreateCategory) (*Category, error) {
	endpoint := fmt.Sprintf("%s/servers/%s/categories", guildedApi, s.client.ServerID)

	var category Category

	err := s.client.PostRequestV2(endpoint, options, &category)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return &category, nil
}

func (s *categoryService) Update(categoryID int, name string) (*Category, error) {
	endpoint := fmt.Sprintf("%s/servers/%s/categories/%d", guildedApi, s.client.ServerID, categoryID)

	var category Category

	body := map[string]interface{}{
		"name": name,
	}

	err := s.client.PatchRequest(endpoint, body, &category)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	return &category, nil
}

func (s *categoryService) Delete(categoryID int) error {
	endpoint := fmt.Sprintf("%s/servers/%s/categories/%d", guildedApi, s.client.ServerID, categoryID)

	_, err := s.client.DeleteRequest(endpoint)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	return nil
}
