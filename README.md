# guildedgo

A guilded.gg library in Go

## Getting started

```cmd
go get github.com/itschip/guildedgo
```

## Examples

### Listen to events

```go
package main

import (
        "fmt"

        "github.com/itschip/guildedgo"
)

func main() {
        guildedClient := guildedgo.NewClient(&guildedgo.Config{
                Token:    "YOUR_TOKEN",
                ServerID: "YOUR_SERVER_ID",
        })

        // Listen to the ChatMessageCreated event
        guildedClient.On("ChatMessageCreated", func(client *guildedgo.Client, v any) {
                data, ok := v.(*guildedgo.ChatMessageCreated)

                if ok {
                        fmt.Println(data.Message.Content)

                        if data.Message.Content == "!ping" {
                                guildedClient.Channel.SendMessage(data.Message.ChannelID, &guildedgo.MessageObject{
                                        Content: "pong!",
                                })
                        }

                }
        })

        // Open socket
        guildedClient.Open()
}
```

### Command builder

```go
serverID := internal.GetEnv("SERVER_ID")
	token := internal.GetEnv("TOKEN")

	config := &guildedgo.Config{
		ServerID: serverID,
		Token:    token,
	}

	c := guildedgo.NewClient(config)

	commands := &guildedgo.CommandsBuilder{
		Commands: []Command{
			{
				CommandName: "!test",
				Action: func(client *Client, v *ChatMessageCreated) {
					client.Channel.SendMessage(v.Message.ChannelID, &MessageObject{
						Content: "Test",
					})

					fmt.Println("Test working")
				},
			},
			{
				CommandName: "!party",
				Action: func(client *Client, v *ChatMessageCreated) {
					client.Channel.SendMessage(v.Message.ChannelID, &MessageObject{
						Content: "Yeah!!! Let's party",
					})

					fmt.Println("Party working")
				},
			},
		},
	}

	c.CommandService.AddCommands(commands)

	c.Open()
```
