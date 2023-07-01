package guildedgo

import (
	"errors"
	"fmt"
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

type ServerResponse struct {
	Server `json:"server"`
}

const (
	ServerTypeTeam         string = "team"
	ServerTypeOrganization string = "organization"
	ServerTypeCommunity    string = "community"
	ServerTypeClan         string = "clan"
	ServerTypeGuild        string = "guild"
	ServerTypeFriends      string = "friends"
	ServerTypeStreaming    string = "streaming"
	ServerTypeOther        string = "other"
)

type ServerService interface {
	GetServer(serverId string) (*Server, error)
}

type serverEndpoints struct{}

func (e *serverEndpoints) Server(serverId string) string {
	return guildedApi + "/servers/" + serverId
}

type serverService struct {
	client    *Client
	endpoints *serverEndpoints
}

var _ ServerService = &serverService{}

func (service *serverService) GetServer(serverId string) (*Server, error) {
	endpoint := service.endpoints.Server(serverId)

	var server ServerResponse
	err := service.client.GetRequestV2(endpoint, &server)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get server. Error: %v", err.Error()))
	}

	return &server.Server, nil
}
