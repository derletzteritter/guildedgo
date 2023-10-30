package user

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

type User struct {
	// The ID of the user
	Id string `json:"id"`

	// The type of user. Can be 'bot' or 'user'. If this property is absent, it can assumed to be of type user
	Type string `json:"type,omitempty"`

	Name string `json:"name"`

	// The avatar image associated with the user
	Avatar string `json:"avatar,omitempty"`

	// The banner image associated with the user
	Banner string `json:"banner,omitempty"`

	// The ISO 8601 timestamp that the user was created at
	CreatedAt string `json:"createdAt"`

	Status Status `json:"status,omitempty"`
}

type Status struct {
	Content string `json:"content,omitempty"`
	EmoteID int    `json:"emoteId"`
}

type UpdateStatusParams struct {
	Content   string `json:"content,omitempty"`
	EmoteID   int    `json:"emoteId"`
	ExpiresAt string `json:"expiresAt,omitempty"`
}

type Summary struct {
	// The ID of the user
	Id string `json:"id"`

	//  The type of user. If this property is absent, it can assumed to be of type user
	Type string `json:"type,omitempty"`

	Name string `json:"name"`

	// The avatar image associated with the user
	Avatar string `json:"avatar,omitempty"`
}

func Get(c Client, userID string) (*User, error) {
	var err error
	endpoint := guildedApi + "/users/" + userID

	var v struct {
		User User `json:"user"`
	}

	res, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("error decoding user: %w", err)
	}

	return &v.User, nil
}

func UpdateStatus(c Client, userID string, params *UpdateStatusParams) error {
	var err error
	endpoint := guildedApi + "/users/" + userID + "/status"

	_, err = c.PerformRequest(http.MethodPut, endpoint, params)
	if err != nil {
		return fmt.Errorf("error updating user status: %v", err)
	}

	return nil
}

func DeleteStatus(c Client, userID string) error {
	var err error
	endpoint := guildedApi + "/users/" + userID + "/status"

	_, err = c.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("error deleting user status: %v", err)
	}

	return nil
}
