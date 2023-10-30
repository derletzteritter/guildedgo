package member

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

func UpdateNickname(c Client, serverID, userId, nickname string) (string, error) {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/members/" + userId + "/nickname"

	var v struct {
		Nickname string `json:"nickname"`
	}

	body := struct {
		Nickname string `json:"nickname"`
	}{
		Nickname: nickname,
	}

	res, err := c.PerformRequest(http.MethodPut, endpoint, body)
	if err != nil {
		return "", fmt.Errorf("failed to update nickname: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return v.Nickname, nil
}

func DeleteNickname(c Client, serverID, userId string) error {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/members/" + userId + "/nickname"

	_, err = c.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to delete nickname: %w", err)
	}

	return nil
}

func Kick(c Client, serverID, userId string) error {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/members/" + userId

	_, err = c.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to kick member: %w", err)
	}

	return nil
}

// Find returns a ServerMember struct for the given user ID
func Find(c Client, serverID, userId string) (*ServerMember, error) {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/members/" + userId

	var v struct {
		Member ServerMember `json:"member"`
	}

	res, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get member: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &v.Member, nil
}

// Get returns all members in the server
func Get(c Client, serverID string) ([]ServerMemberSummary, error) {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/members"

	var v struct {
		Members []ServerMemberSummary `json:"members"`
	}

	res, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get members: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return v.Members, nil
}
