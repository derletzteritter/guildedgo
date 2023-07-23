package guildedgo

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/itschip/guildedgo/internal"
)

func TestNewClient(t *testing.T) {
	serverID := internal.GetEnv("SERVER_ID")
	token := internal.GetEnv("TOKEN")

	config := &Config{
		ServerID: serverID,
		Token:    token,
	}

	c := NewClient(config)

	c.Command("ping", func(client *Client, v *ChatMessageCreated) {
		client.Channel.SendMessage(v.Message.ChannelID, &MessageObject{
			Content: "Pong!",
		})
	})

	c.Open()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt

	c.Close()
}
