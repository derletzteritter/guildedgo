package client_test

import (
	"os"
	"testing"

	"github.com/itschip/guildedgo/pkg/client"
	"github.com/itschip/guildedgo/pkg/message"
	"github.com/joho/godotenv"
)

func TestSendMessage(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	token := os.Getenv("TOKEN")
	serverID := os.Getenv("SERVER_ID")

	c := client.New(client.Config{
		Token:    token,
		ServerID: serverID,
	})

	msg, err := message.Send(c, "", message.MessageParams{
		Content: "Hello, world!",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log("Message content", msg.Content)
}

func TestGetMessages(t *testing.T) {
	t.Skip("Skipping test as api is broken")

	err := godotenv.Load("../../.env")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	token := os.Getenv("TOKEN")
	serverID := os.Getenv("SERVER_ID")

	c := client.New(client.Config{
		Token:    token,
		ServerID: serverID,
	})

	msg, err := message.Get(c, "", "")
	if err != nil {
		t.Error(err)
	}

	t.Log(msg)
}
