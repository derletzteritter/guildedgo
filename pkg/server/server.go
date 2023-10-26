package server

import (
	"fmt"
	"net/http"

	"github.com/itschip/guildedgo/pkg/client"
)

type Server struct {
	ID         string `json:"id"`
	OwnerID    string `json:"ownerid"`
	Type       string `json:"type,omitempty"`
	Name       string `json:"name"`
	URL        string `json:"url,omitempty"`
	About      string `json:"about,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	Banner     string `json:"banner,omitempty"`
	Timezone   string `json:"timezone,omitempty"`
	IsVerified bool   `json:"isVerified,omitempty"`

	// The channel ID of the default channel of the server.
	// This channel is defined as the first chat or voice channel in the left sidebar of a server in our UI.
	// This channel is useful for sending welcome messages,
	// though note that a bot may not have permissions to interact with this channel depending on how the server is configured.
	DefaultChannelId string `json:"defaultChannelId,omitempty"`
	CreatedAt        string `json:"createdAt"`
}

type serverResponse struct {
	Server `json:"server"`
}

const (
	TypeTeam         string = "team"
	TypeOrganization string = "organization"
	TypeCommunity    string = "community"
	TypeClan         string = "clan"
	TypeGuild        string = "guild"
	TypeFriends      string = "friends"
	TypeStreaming    string = "streaming"
	TypeOther        string = "other"
)

func Get(c *client.Client) (Server, error) {
	var err error
	endpoint := client.GuildedApi + "/servers/" + c.ServerID

	var v serverResponse

	body, err := c.Http.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return Server{}, fmt.Errorf("failed to get server: %w", err)
	}

	err = c.Http.Decode(body, &v)
	if err != nil {
		return Server{}, fmt.Errorf("failed to decode server response: %w", err)
	}

	return v.Server, nil
}
