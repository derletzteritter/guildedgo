package category_test

import (
	"github.com/itschip/guildedgo/pkg/category"
	"github.com/itschip/guildedgo/pkg/client"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
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

	cat, err := category.Create(c, &category.CreateParams{
		Name: "Test Category",
	})

	if err != nil {
		t.Error(err)
	}

	t.Log("Created category")
	t.Log(cat)
}
