package membership

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Client interface {
	PerformRequest(method, url string, data any) (io.ReadCloser, error)
	Decode(body io.ReadCloser, v any) error
}

const (
	guildedApi = "https://www.guilded.gg/api/v1"
)

func AddGroupMember(c Client, groupID, userID string) error {
	endpoint := guildedApi + "/groups/" + groupID + "/members/" + userID

	_, err := c.PerformRequest(http.MethodPut, endpoint, nil)
	if err != nil {
		return fmt.Errorf("error adding member to group: %w", err)
	}

	return nil
}

func RemoveGroupMember(c Client, groupID, userID string) error {
	endpoint := guildedApi + "/groups/" + groupID + "/members/" + userID

	_, err := c.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("error removing member from group: %w", err)
	}

	return nil
}

func AssignRole(c Client, serverID, userID string, roleID int) error {
	endpoint := guildedApi + "/servers/" + serverID + "/members/" + userID + "/roles/" + strconv.Itoa(roleID)

	_, err := c.PerformRequest(http.MethodPut, endpoint, nil)
	if err != nil {
		return fmt.Errorf("error assigning role to member: %w", err)
	}

	return nil
}

func RemoveRole(c Client, serverID, userID string, roleID int) error {
	endpoint := guildedApi + "/servers/" + serverID + "/members/" + userID + "/roles/" + strconv.Itoa(roleID)

	_, err := c.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("error removing role from member: %w", err)
	}

	return nil
}

func GetRoles(c Client, serverID, userID string) ([]int, error) {
	var err error
	endpoint := guildedApi + "/servers/" + serverID + "/members/" + userID + "/roles"

	var v []int

	res, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting roles: %w", err)
	}

	err = c.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("error decoding roles: %w", err)
	}

	return v, nil
}
