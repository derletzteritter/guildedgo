package client

import (
	"github.com/gorilla/websocket"
	"github.com/itschip/guildedgo/internal/http"
	"os"
	"sync"
)

type Client struct {
	sync.RWMutex
	wsMutex   sync.Mutex
	conn      *websocket.Conn
	interrupt chan os.Signal
	listening chan struct{}
	events    map[string][]event

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
