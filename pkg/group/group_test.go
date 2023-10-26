package group_test

import (
	"os"
	"testing"

	"github.com/itschip/guildedgo/pkg/client"
	"github.com/itschip/guildedgo/pkg/group"
	"github.com/joho/godotenv"
)

func TestGetGroups(t *testing.T) {
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

	groups, err := group.Get(c)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(groups)
}
