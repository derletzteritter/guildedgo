package member

import (
	"fmt"
	"github.com/itschip/guildedgo/pkg/client"
	"github.com/itschip/guildedgo/pkg/user"
	"net/http"
)

type ServerMember struct {
	User user.User `json:"user"`

	// (must have unique items true)
	RoleIds []int `json:"roleIds"`

	Nickname string `json:"nickname,omitempty"`

	// The ISO 8601 timestamp that the member was created at
	JoinedAt string `json:"joinedAt"`

	// (default false)
	IsOwner bool `json:"isOwner,omitempty"`
}

type ServerMemberSummary struct {
	User user.Summary `json:"user"`

	// (must have unique items true)
	RoleIds []int `json:"roleIds"`
}

type ServerMemberPermissions struct {
	Permissions []string `json:"permissions"`
}

func UpdateNickname(c *client.Client, userId, nickname string) (string, error) {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/members/" + userId + "/nickname"

	var v struct {
		Nickname string `json:"nickname"`
	}

	body := struct {
		Nickname string `json:"nickname"`
	}{
		Nickname: nickname,
	}

	res, err := c.Http.PerformRequest(http.MethodPut, endpoint, body)
	if err != nil {
		return "", fmt.Errorf("failed to update nickname: %w", err)
	}

	err = c.Http.Decode(res, &v)
	if err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return v.Nickname, nil
}

func DeleteNickname(c *client.Client, userId string) error {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/members/" + userId + "/nickname"

	_, err = c.Http.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to delete nickname: %w", err)
	}

	return nil
}

func Kick(c *client.Client, userId string) error {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/members/" + userId

	_, err = c.Http.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to kick member: %w", err)
	}

	return nil
}

// Find returns a ServerMember struct for the given user ID
func Find(c *client.Client, userId string) (*ServerMember, error) {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/members/" + userId

	var v struct {
		Member ServerMember `json:"member"`
	}

	res, err := c.Http.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get member: %w", err)
	}

	err = c.Http.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Member, nil
}

// Get returns all members in the server
func Get(c *client.Client) ([]ServerMemberSummary, error) {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/members"

	var v struct {
		Members []ServerMemberSummary `json:"members"`
	}

	res, err := c.Http.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get members: %w", err)
	}

	err = c.Http.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return v.Members, nil
}
