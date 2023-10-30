package social

type Link struct {
	Type      string `json:"type"`
	UserID    string `json:"userId"`
	Handle    string `json:"handle,omitempty"`
	ServiceID string `json:"serviceId,omitempty"`
	CreatedAt string `json:"createdAt"`
}
