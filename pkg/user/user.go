package user

type User struct {
	// The ID of the user
	Id string `json:"id"`

	// The type of user. Can be 'bot' or 'user'. If this property is absent, it can assumed to be of type user
	Type string `json:"type,omitempty"`

	Name string `json:"name"`

	// The avatar image associated with the user
	Avatar string `json:"avatar,omitempty"`

	// The banner image associated with the user
	Banner string `json:"banner,omitempty"`

	// The ISO 8601 timestamp that the user was created at
	CreatedAt string `json:"createdAt"`

	Status Status `json:"status,omitempty"`
}

type Status struct {
	Content string `json:"content,omitempty"`
	EmoteID int    `json:"emoteId"`
}

type Summary struct {
	// The ID of the user
	Id string `json:"id"`

	//  The type of user. If this property is absent, it can assumed to be of type user
	Type string `json:"type,omitempty"`

	Name string `json:"name"`

	// The avatar image associated with the user
	Avatar string `json:"avatar,omitempty"`
}
