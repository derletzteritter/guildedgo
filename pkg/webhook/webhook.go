package webhook

type Webhook struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar,omitempty"`
	ServerID  string `json:"serverId"`
	ChannelID string `json:"channelId"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	DeletedAt string `json:"deletedAt,omitempty"`
	Token     string `json:"token,omitempty"`
}
