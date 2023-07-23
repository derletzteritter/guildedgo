package channel

import "github.com/itschip/guildedgo/pkg/client"

type Service struct {
	client *client.Client
}

func NewService(client *client.Client) Service {
	return Service{
		client,
	}
}
