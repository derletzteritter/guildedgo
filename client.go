package guildedgo

import (
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	guildedApi = "https://www.guilded.gg/api/v1"
)

type Client struct {
	sync.RWMutex
	wsMutex        sync.Mutex
	Token          string
	ServerID       string
	client         *http.Client
	conn           *websocket.Conn
	interrupt      chan os.Signal
	listening      chan struct{}
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
	DocComments    DocCommentService
	Docs           DocsService
	Socials        SocialsService
	Announcements  AnnouncementService
	Category       CategoryService
	Users          UserService
	events         map[string][]Event
	commands       map[string]Command
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
	c.Docs = &docsService{client: c}
	c.DocComments = &docCommentService{client: c}
	c.Socials = &socialsService{client: c}
	c.Announcements = &announcementService{client: c}
	c.Users = &userService{client: c}
	c.Category = &categoryService{client: c}

	c.events = make(map[string][]Event)

	return c
}
