package membership

import (
	"fmt"
	"github.com/itschip/guildedgo/pkg/client"
	"net/http"
	"strconv"
)

func AddGroupMember(c *client.Client, groupID, userID string) error {
	endpoint := client.GuildedApi + "/groups/" + groupID + "/members/" + userID

	_, err := c.Http.PerformRequest(http.MethodPut, endpoint, nil)
	if err != nil {
		return fmt.Errorf("error adding member to group: %w", err)
	}

	return nil
}

func RemoveGroupMember(c *client.Client, groupID, userID string) error {
	endpoint := client.GuildedApi + "/groups/" + groupID + "/members/" + userID

	_, err := c.Http.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("error removing member from group: %w", err)
	}

	return nil
}

func AssignRole(c *client.Client, userID string, roleID int) error {
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/members/" + userID + "/roles/" + strconv.Itoa(roleID)

	_, err := c.Http.PerformRequest(http.MethodPut, endpoint, nil)
	if err != nil {
		return fmt.Errorf("error assigning role to member: %w", err)
	}

	return nil
}

func RemoveRole(c *client.Client, userID string, roleID int) error {
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/members/" + userID + "/roles/" + strconv.Itoa(roleID)

	_, err := c.Http.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("error removing role from member: %w", err)
	}

	return nil
}

func GetRoles(c *client.Client, userID string) ([]int, error) {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID + "/members/" + userID + "/roles"

	var v []int

	res, err := c.Http.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting roles: %w", err)
	}

	err = c.Http.Decode(res, &v)
	if err != nil {
		return nil, fmt.Errorf("error decoding roles: %w", err)
	}

	return v, nil
}
