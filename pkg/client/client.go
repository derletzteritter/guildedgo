package client

import (
	"github.com/gorilla/websocket"
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
	Token    string
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
		Token:    config.Token,
		events:   make(map[string][]event),
	}
}
