package ban

import (
	"fmt"
	"github.com/itschip/guildedgo/pkg/user"
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

// Ban is the ServerMemberBan model
type Ban struct {
	User   user.Summary `json:"user"`
	Reason string       `json:"reason"`
	// CreatedBy is the ID of the user who banned the user
	CreatedBy string `json:"createdBy"`
	CreatedAt string `json:"createdAt"`
}

type CreateParams struct {
	Reason string `json:"reason"`
}

func Create(c Client, serverID, userID string, params *CreateParams) (*Ban, error) {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/bans/" + userID

	var v struct {
		Ban `json:"serverMemberBan"`
	}

	res, err := c.PerformRequest(http.MethodPost, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("failed to ban member: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Ban, nil
}

// Find returns the ban for the given user ID
func Find(c Client, serverID, userID string) (*Ban, error) {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/bans/" + userID

	var v struct {
		Ban `json:"serverMemberBan"`
	}

	res, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get ban: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Ban, nil
}

// Delete deletes the ban for the given user ID
func Delete(c Client, serverID, userID string) error {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/bans/" + userID

	res, err := c.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to delete ban: %w", err)
	}

	err = c.Decode(res, nil)
	if err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

// Get returns a list of bans for the server
func Get(c Client, serverID string) ([]Ban, error) {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/bans"

	var v struct {
		Bans []Ban `json:"serverMemberBan"`
	}

	res, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get bans: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return v.Bans, nil
}
