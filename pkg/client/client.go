package client

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	Token    string
	ServerID string
	client   *http.Client
	conn     *websocket.Conn
}
