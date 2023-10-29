package member_test

import (
	"github.com/itschip/guildedgo/pkg/client"
	"github.com/itschip/guildedgo/pkg/member"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestUpdateNickname(t *testing.T) {
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

	nickname, err := member.UpdateNickname(c, "", "funnyguy123")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(nickname)
}
