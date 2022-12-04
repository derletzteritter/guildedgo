package guildedgo

// ChatMessage is the default message struct
type ChatMessage struct {
	// The ID of the message
	ID string `json:"id"`

	// The type of chat message.
	// "system" messages are generated by Guilded, while "default" messages are user or bot-generated.
	Type string `json:"type,omitempty"`

	// The ID of the server
	ServerID string `json:"serverId"`

	// The ID of the channel
	ChannelID string `json:"channelId"`

	// The content of the message
	Content string `json:"content,omitempty"`

	// Embeds for the message.
	//
	// (min items 1; max items 10
	Embeds []ChatEmbed `json:"embeds,omitempty"`

	// Message IDs that were replied to
	ReplyMessageIds []string `json:"replyMessageIds,omitempty"`

	// If set, this message will only be seen by those mentioned or replied to.
	IsPrivate bool `json:"isPrivate,omitempty"`

	// If set, this message did not notify mention or reply recipients (default false)
	IsSilent bool `json:"isSilent,omitempty"`

	Mentions `json:"mentions,omitempty"`

	// The ISO 8601 timestamp that the message was created at.
	CreatedAt string `json:"createdAt"`

	// The ID for the user who created this messsage
	// (Note: if this event has `createdByWebhookId` present,
	// this field will still be populated, but can be ignored.
	// In this case, the value of this field will always be Ann6LewA)
	CreatedBy string `json:"createdBy"`

	// The ID of the webhook who created this message, if was created by a webhook
	CreatedByWebhookId string `json:"createdByWebhookId"`

	// The IOSO 8601 timestamp that the message was updated at, if relevant
	UpdatedAt string `json:"updatedAt"`
}

// ChatEmbed are rich content sections optionally associated with chat messages.
// Properties with "webhook-markdown" support allow for the following: link, italic,
// bold, strikethrough, underline, inline code, block code, reaction, and mention.
type ChatEmbed struct {
	// Main header of the embed (max length 256)
	Title string `json:"embed,omitempty"`

	// Subtext of the embed (max length 2048)
	Description string `json:"description,omitempty"`

	// URL to linkify the title field with (max length 1024; regex ^(?!attachment))
	URL string `json:"url,omitempty"`

	// Decimal value of the color that the left border should be (min 0; max 16777215)
	Color int `json:"color,omitempty"`

	// A small section at the bottom of the embed
	Footer ChatEmbedFooter `json:"footer,omitempty"`

	// A timestamp to put in the footer
	Timestamp string `json:"timestamp,omitempty"`

	// An image to the right of the embed's content
	Thumbnail ChatEmbedThumbnail `json:"thumbnail,omitempty"`

	// The main picture to associate with the embed
	Image ChatEmbedImage `json:"image"`

	// A small section above the title of the embed
	Author ChatEmbedAuthor `json:"author,omitempty"`

	// Table-like cells to add to the embed (max items 25)
	Fields []ChatEmbedField `json:"fields,omitempty"`
}

type ChatEmbedFooter struct {
	// URL of a small image to put in the footer (max length 1024)
	IconURL string `json:"icon_url,omitempty"`

	// Text of the footer (max length 2048)
	Text string `json:"text"`
}

type ChatEmbedThumbnail struct {
	// URL of the image (max length 1024)
	URL string `json:"url,omitempty"`
}

type ChatEmbedImage struct {
	// URL of the image (max length 1024)
	URL string `json:"url,omitempty"`
}

type ChatEmbedAuthor struct {
	// Name of the author (max length 256)
	Name string `jsojn:"name,omitempty"`

	// URL to linkify the author's name field (max length 1024; regex ^(?!attachment))
	URL string `json:"url,omitempty"`

	// URL of a small image to display to the left of the author's name (max length 1024
	IconURL string `json:"icon_url,omitempty"`
}

type ChatEmbedField struct {
	// Header of the table-like cell (max length 256)
	Name string `json:"name"`

	// Subtext of the table-like cell (max length 1024)
	Value string `json:"value"`

	// If the field should wrap or not (default false)
	Inline bool `json:"inline"`
}

type MessageObject struct {
	// If set, this message will only be seen by those mentioned or replied to
	IsPrivate string `json:"isPrivate,omitempty"`

	// If set, this message will not notify any mentioned users or roles (default false)
	IsSilent string `json:"isSilent,omitempty"`

	// Message IDs to reply to (min items 1; max items 5)
	ReplyMessageIds []string `json:"replyMessageIds,omitempty"`

	// The content of the message (min length 1; max length 4000)
	Content string `json:"content,omitempty"`

	// At this time, only one embed is supported per message, and attachments are not supported.
	// If you need to send more than one embed or upload attachments, consider creating the
	// message via a webhook. (min items 1; max items 1)
	Embeds []ChatEmbed `json:"embeds,omitempty"`
}

type MessageResponse struct {
	Message ChatMessage `json:"message"`
}

type GetMessageResponse struct {
	Message ChatMessage `json:"message"`
}

type GetMessagesObject struct {
	Before         string `json:"before"`
	After          string `json:"after"`
	Limit          int    `json:"limit"`
	IncludePrivate bool   `json:"includePrivate"`
}

type AllMessagesResponse struct {
	Messages []ChatMessage `json:"messages"`
}
