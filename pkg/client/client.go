package client

import (
	"github.com/itschip/guildedgo/internal/http"
)

type Client struct {
	ServerID string
	Http     *http.Http
}

type Config struct {
	Token    string
	ServerID string
}

const (
	GuildedApi = "https://www.guilded.gg/api/v1"
)

func New(config Config) *Client {
	return &Client{
		ServerID: config.ServerID,
		Http: &http.Http{
			Token: config.Token,
		},
	}
}
