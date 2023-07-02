package guildedgo

import (
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

	c.Announcements.GetAnnouncements("123", &GetAnnouncementParams{
		Before: "123",
		Limit:  1,
	})
}
