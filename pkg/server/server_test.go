package server_test

import (
	"github.com/itschip/guildedgo/pkg/client"
	"os"
	"testing"

	"github.com/itschip/guildedgo/pkg/server"
	"github.com/joho/godotenv"
)

func TestGetServer(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	token := os.Getenv("TOKEN")
	serverID := os.Getenv("SERVER_ID")

	c := client.New(client.Config{
		ServerID: serverID,
		Token:    token,
	})

	s, err := server.Get(c, serverID)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(s)
}
