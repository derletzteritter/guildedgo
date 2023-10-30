package event

import (
	"github.com/itschip/guildedgo/pkg/ban"
	"github.com/itschip/guildedgo/pkg/channel"
	"github.com/itschip/guildedgo/pkg/member"
	"github.com/itschip/guildedgo/pkg/message"
	"github.com/itschip/guildedgo/pkg/server"
	"github.com/itschip/guildedgo/pkg/social"
	"github.com/itschip/guildedgo/pkg/webhook"
)

type BotServerMembershipCreated struct {
	Server    server.Server `json:"server"`
	CreatedBy string        `json:"createdBy"`
}

type BotServerMembershipDeleted struct {
	Server    server.Server `json:"server"`
	DeletedBy string        `json:"deletedBy"`
}

type ChatMessageCreated struct {
	ServerID string              `json:"serverId"`
	Message  message.ChatMessage `json:"message"`
}

type ChatMessageUpdated struct {
	ServerID string              `json:"serverId"`
	Message  message.ChatMessage `json:"message"`
}

type ChatMessageDeleted struct {
	ServerID  string `json:"serverId"`
	DeletedAt string `json:"deletedAt"`
	Message   struct {
		ID                    string              `json:"id"`
		Type                  string              `json:"type"`
		ServerID              string              `json:"serverId,omitempty"`
		GroupID               string              `json:"groupId,omitempty"`
		ChannelID             string              `json:"channelId"`
		Content               string              `json:"content,omitempty"`
		HiddenLinkPreviewUrls []string            `json:"hiddenLinkPreviewUrls,omitempty"`
		Embeds                []message.ChatEmbed `json:"embeds,omitempty"`
		ReplyMessageIds       []string            `json:"replyMessageIds,omitempty"`
		IsPrivate             bool                `json:"isPrivate,omitempty"`
		IsSilent              bool                `json:"isSilent,omitempty"`
		IsPinned              bool                `json:"isPinned,omitempty"`
		Mentions              channel.Mention     `json:"mentions,omitempty"`
		CreatedAt             string              `json:"createdAt"`
		CreatedBy             string              `json:"createdBy"`
		CreatedByWebhookID    string              `json:"createdByWebhookId,omitempty"`
		UpdatedAt             string              `json:"updatedAt,omitempty"`
		DeletedAt             string              `json:"deletedAt"`
	} `json:"message"`
}

type ServerMemberJoined struct {
	ServerID          string              `json:"serverId"`
	Member            member.ServerMember `json:"member"`
	ServerMemberCount int                 `json:"serverMemberCount"`
}

type ServerMemberRemoved struct {
	ServerID string `json:"serverId"`
	UserID   string `json:"userId"`
	IsKick   bool   `json:"isKick,omitempty"`
	IsBan    bool   `json:"isBan,omitempty"`
}

type ServerMemberBanned struct {
	ServerID string  `json:"serverId"`
	Ban      ban.Ban `json:"serverMemberBan"`
}

type ServerMemberUnbanned struct {
	ServerID string  `json:"serverId"`
	Ban      ban.Ban `json:"serverMemberBan"`
}

type ServerMemberUpdated struct {
	ServerID string `json:"serverId"`
	UserInfo struct {
		ID       string `json:"id"`
		Nickname string `json:"nickname,omitempty"`
	} `json:"userInfo"`
}

type ServerRolesUpdated struct {
	ServerID      string `json:"serverId"`
	MemberRoleIds []struct {
		UserID  string   `json:"userId"`
		RoleIds []string `json:"roleIds"`
	} `json:"memberRoleIds"`
}

type ServerChannelCreated struct {
	ServerID string                `json:"serverId"`
	Channel  channel.ServerChannel `json:"channel"`
}

type ServerChannelUpdated struct {
	ServerID string                `json:"serverId"`
	Channel  channel.ServerChannel `json:"channel"`
}

type ServerChannelDeleted struct {
	ServerID string                `json:"serverId"`
	Channel  channel.ServerChannel `json:"channel"`
}

type ServerMemberSocialLinkCreated struct {
	ServerID   string      `json:"serverId"`
	SocialLink social.Link `json:"socialLink"`
}

type ServerMemberSocialLinkUpdated struct {
	ServerID   string      `json:"serverId"`
	SocialLink social.Link `json:"socialLink"`
}

type ServerMemberSocialLinkDeleted struct {
	ServerID   string      `json:"serverId"`
	SocialLink social.Link `json:"socialLink"`
}

type ServerWebhookCreated struct {
	ServerID string          `json:"serverId"`
	Webhook  webhook.Webhook `json:"webhook"`
}

type ServerWebhookUpdated struct {
	ServerID string          `json:"serverId"`
	Webhook  webhook.Webhook `json:"webhook"`
}
type ChannelArchived struct {
	ServerID string                `json:"serverId"`
	Channel  channel.ServerChannel `json:"channel"`
}

type ChannelRestored struct {
	ServerID string                `json:"serverId"`
	Channel  channel.ServerChannel `json:"channel"`
}
