package guildedgo

type Webhook struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ServerID  string `json:"serverId"`
	ChannelID string `json:"channelId"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	DeletedAt string `json:"deletedAt"`
	Token     string `json:"string"`
}

type BotServerMembershipCreated struct {
	Server `json:"server"`

	// The ID of the user who created this server membership
	CreatedBy string `json:"createdBy"`
}

type BotServerMembershipDeleted struct {
	Server `json:"server"`

	// The ID of the user who deleted this server membership
	CreatedBy string `json:"createdBy"`
}

type ChatMessageCreated struct {
	// The ID of the server
	ServerID string `json:"serverId"`

	Message ChatMessage `json:"message"`
}

type ChatMessageUpdated struct {
	// The ID of the server
	ServerID string `json:"serverId"`

	Message ChatMessage `json:"message"`
}

type ChatMessageDeleted struct {
	// The ID of the server
	ServerID string `json:"serverId"`

	Message struct {
		// The ID of the message
		ID string `json:"id"`

		// The ID of the serve
		ServerID string `json:"serverId,omitempty"`

		// The ID of the channel
		ChannelID string `json:"channelId"`

		DeletedAt string `json:"deletedAt"`

		IsPrivate bool `json:"isPrivate,omitempty"`
	} `json:"message"`
}

type ServerMemberJoined struct {
	// The ID of the server
	ServerID string `json:"serverId"`

	Member ServerMember `json:"member"`
}

type ServerMemberRemoved struct {
	// The ID of the server
	ServerID string `json:"serverId"`

	// The ID of the user
	UserID string `json:"userId"`

	// If this member leaving was the result of a kick
	IsKick bool `json:"isKick,omitempty"`

	// If this member leaving was the result of a ban
	IsBan bool `json:"isBan,omitempty"`
}

type ServerMemberBanned struct {
	// The ID of the server
	ServerID string `json:"serverId"`

	ServerMemberBan `json:"serverMemberBan"`
}

type ServerMemberUnbanned struct {
	// The ID of the server
	ServerID string `json:"serverId"`

	ServerMemberBan `json:"serverMemberBan"`
}

type ServerMemberUpdated struct {
	// The ID of the server
	ServerID string `json:"serverId"`

	UserInfo struct {
		// The ID of the user
		ID string `json:"userId"`

		// The nickname that was just updated for the user
		Nickname string `json:"nickname,omitempty"`
	} `json:"userInfo"`
}

type ServerRolesUpdated struct {
	// The ID of the server
	ServerID string `json:"serverId"`

	MemberRoleIds struct {
		UserID  string `json:"userId"`
		RoleIDs []int  `json:"roleIds"`
	} `json:"memberRoleIds"`
}

type ServerChannelCreated struct {
	ServerID string        `json:"serverId"`
	Channel  ServerChannel `json:"channel"`
}

type ServerChannelUpdated struct {
	ServerID string        `json:"serverId"`
	Channel  ServerChannel `json:"channel"`
}

type ServerChannelDeleted struct {
	ServerID string        `json:"serverId"`
	Channel  ServerChannel `json:"channel"`
}

type ServerWebhookCreated struct {
	ServerID string `json:"serverId"`
	Webhook  `json:"webhook"`
}

type ServerWebhookUpdated struct {
	ServerID string `json:"serverId"`
	Webhook  `json:"webhook"`
}

type DocCreated struct {
	ServerID string `json:"serverId"`
	Doc      `json:"doc"`
}

type DocUpdated struct {
	ServerID string `json:"serverId"`
	Doc      `json:"doc"`
}

type DocDeleted struct {
	ServerID string `json:"serverId"`
	Doc      `json:"doc"`
}

type CalendarEventCreated struct {
	ServerID      string `json:"serverId"`
	CalendarEvent `json:"calendarEvent"`
}

type CalendarEventUpdated struct {
	ServerID      string `json:"serverId"`
	CalendarEvent `json:"calendarEvent"`
}

type CalendarEventDeleted struct {
	ServerID      string `json:"serverId"`
	CalendarEvent `json:"calendarEvent"`
}

type ForumTopicCreated struct {
	ServerID   string `json:"serverId"`
	ForumTopic `json:"forumTopic"`
}

type ForumTopicUpdated struct {
	ServerID   string `json:"serverId"`
	ForumTopic `json:"forumTopic"`
}

type ForumTopicDeleted struct {
	ServerID   string `json:"serverId"`
	ForumTopic `json:"forumTopic"`
}

type ForumTopicPinned struct {
	ServerID   string `json:"serverId"`
	ForumTopic `json:"forumTopic"`
}

type ForumTopicUnpinned struct {
	ServerID   string `json:"serverId"`
	ForumTopic `json:"forumTopic"`
}

type ForumTopicReactionCreated struct {
	ServerID string `json:"serverId"`
	Reaction struct {
		ChannelID    string `json:"channelId"`
		CreatedBy    string `json:"createdBy"`
		Emote        `json:"emote"`
		ForumTopicID int `json:"forumTopicId"`
	}
}

type ForumTopicReactionDeleted struct {
	ServerID string `json:"serverId"`
	Reaction struct {
		ChannelID    string `json:"channelId"`
		CreatedBy    string `json:"createdBy"`
		Emote        `json:"emote"`
		ForumTopicID int `json:"forumTopicId"`
	}
}

type ForumTopicLocked struct {
	ServerID   string `json:"serverId"`
	ForumTopic `json:"forumTopic"`
}

type ForumTopicUnlocked struct {
	ServerID   string `json:"serverId"`
	ForumTopic `json:"forumTopic"`
}

type ForumTopicCommentCreated struct {
	ServerID          string `json:"serverId"`
	ForumTopicComment `json:"forumTopicComment"`
}

type ForumTopicCommentUpdated struct {
	ServerID          string `json:"serverId"`
	ForumTopicComment `json:"forumTopicComment"`
}

type ForumTopicCommentDeleted struct {
	ServerID          string `json:"serverId"`
	ForumTopicComment `json:"forumTopicComment"`
}

type ChannelArchived struct {
	ServerID string        `json:"serverId"`
	Channel  ServerChannel `json:"channel"`
}

type ChannelRestored struct {
	ServerID string        `json:"serverId"`
	Channel  ServerChannel `json:"channel"`
}
