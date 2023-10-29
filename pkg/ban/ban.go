package ban

import (
	"fmt"
	"github.com/itschip/guildedgo/pkg/client"
	"github.com/itschip/guildedgo/pkg/user"
	"net/http"
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

func Create(c *client.Client, userID string, params *CreateParams) (*Ban, error) {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/bans/" + userID

	var v struct {
		Ban `json:"serverMemberBan"`
	}

	res, err := c.Http.PerformRequest(http.MethodPost, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("failed to ban member: %w", err)
	}

	err = c.Http.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Ban, nil
}

// Find returns the ban for the given user ID
func Find(c *client.Client, userID string) (*Ban, error) {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/bans/" + userID

	var v struct {
		Ban `json:"serverMemberBan"`
	}

	res, err := c.Http.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get ban: %w", err)
	}

	err = c.Http.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Ban, nil
}

// Delete deletes the ban for the given user ID
func Delete(c *client.Client, userID string) error {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/bans/" + userID

	res, err := c.Http.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to delete ban: %w", err)
	}

	err = c.Http.Decode(res, nil)
	if err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

// Get returns a list of bans for the server
func Get(c *client.Client) ([]Ban, error) {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/bans"

	var v struct {
		Bans []Ban `json:"serverMemberBan"`
	}

	res, err := c.Http.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get bans: %w", err)
	}

	err = c.Http.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return v.Bans, nil
}
