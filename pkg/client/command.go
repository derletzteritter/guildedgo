package client

import socketevent "github.com/itschip/guildedgo/pkg/event"

type CommandsBuilder struct {
	Commands []Command
}

type Command struct {
	CommandName string
	Action      func(client *Client, v *socketevent.ChatMessageCreated)
}

func (r *Client) AddCommands(builder *CommandsBuilder) {
	// Is this the best way to do this? I'm not sure. - Thanks, Copilot
	for _, command := range builder.Commands {
		r.Command(command.CommandName, command.Action)
	}
}

func (r *Client) Command(cmd string, callback func(client *Client, v *socketevent.ChatMessageCreated)) {
	r.On("ChatMessageCreated", func(client *Client, v any) {
		data, ok := v.(*socketevent.ChatMessageCreated)
		if ok {
			if data.Message.Content == cmd {
				callback(client, data)
			}
		}
	})
}
