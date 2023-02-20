package guildedgo

import "net/http"

const (
	guildedApi = "https://www.guilded.gg/api/v1"
)

type Client struct {
	Token    string
	ServerID string
	client   *http.Client

	Channel        ChannelService
	Members        MembersService
	Roles          RoleService
	Server         ServerService
	Forums         ForumService
	Calendar       CalendarService
	Reactions      ReactionService
	List           ListService
	Webhooks       WebhookService
	ServerXP       ServerXPService
	CommandService CommandService
	Events         map[string]Event
}

type Event struct {
	Callback func(*Client, any)
	Type     *interface{}
}

type Config struct {
	Token    string
	ServerID string
}

func NewClient(config *Config) *Client {
	c := &Client{
		Token:    config.Token,
		ServerID: config.ServerID,
		client:   http.DefaultClient,
	}

	c.Channel = &channelService{client: c}
	c.Members = &membersService{client: c}
	c.Roles = &roleService{client: c}
	c.Server = &serverService{client: c}
	c.Forums = &forumService{client: c}
	c.Calendar = &calendarService{client: c}
	c.Reactions = &reactionService{client: c}
	c.CommandService = &commandService{client: c}
	c.List = &listService{client: c}
	c.Webhooks = &webhookService{client: c}
	c.ServerXP = &serverXPService{client: c}

	c.Events = make(map[string]Event)

	return c
}
