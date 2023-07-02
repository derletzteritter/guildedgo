package guildedgo

type User struct {
	// The ID of the user
	Id string `json:"id"`

	// The type of user. If this property is absent, it can assumed to be of type user
	Type string `json:"type,omitempty"`

	Name string `json:"name"`

	// The avatar image associated with the user
	Avatar string `json:"avatar,omitempty"`

	// The banner image associated with the user
	Banner string `json:"banner,omitempty"`

	// The ISO 8601 timestamp that the user was created at
	CreatedAt string `json:"createdAt"`

	Status UserStatus `json:"status,omitempty"`
}

type UserStatus struct {
	Content string `json:"content,omitempty"`
	EmoteID int    `json:"emoteId"`
}

type UserSummary struct {
	// The ID of the user
	Id string `json:"id"`

	//  The type of user. If this property is absent, it can assumed to be of type user
	Type string `json:"type,omitempty"`

	Name string `json:"name"`

	// The avatar image associated with the user
	Avatar string `json:"avatar,omitempty"`
}

type UserStatusUpdate struct {
	Content   string `json:"content,omitempty"`
	EmoteID   int    `json:"emoteId"`
	ExpiresAt string `json:"expiresAt,omitempty"`
}

type UserResponse struct {
	User `json:"user"`
}

const (
	UserTypeUser = "user"
	UserTypeBot  = "bot"
)

type UserService interface {
	GetUser(id string) (*User, error)
	GetUsersServers(id string) ([]Server, error)
	UpdateStatus(id string, status UserStatusUpdate) error
	DeleteStatus(id string) error
}

type userService struct {
	client *Client
}

var _ UserService = &userService{}

func (service *userService) GetUser(id string) (*User, error) {
	endpoint := guildedApi + "/users/" + id

	var response UserResponse
	err := service.client.GetRequestV2(endpoint, &response)
	if err != nil {
		return nil, err
	}
	return &response.User, nil
}

func (service *userService) GetUsersServers(id string) ([]Server, error) {
	endpoint := guildedApi + "/users/" + id + "/servers"

	var response struct {
		Servers []Server `json:"servers"`
	}
	err := service.client.GetRequestV2(endpoint, &response)
	if err != nil {
		return nil, err
	}
	return response.Servers, nil
}

func (service *userService) UpdateStatus(id string, status UserStatusUpdate) error {
	endpoint := guildedApi + "/users/" + id + "/status"

	err := service.client.PutRequestV2(endpoint, &status, nil)
	if err != nil {
		return err
	}

	return nil
}

func (service *userService) DeleteStatus(id string) error {
	endpoint := guildedApi + "/users/" + id + "/status"

	_, err := service.client.DeleteRequest(endpoint)
	if err != nil {
		return err
	}

	return nil
}
