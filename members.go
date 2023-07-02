package guildedgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type ServerMember struct {
	User User `json:"user"`

	// (must have unique items true)
	RoleIds []int `json:"roleIds"`

	Nickname string `json:"nickname,omitempty"`

	// The ISO 8601 timestamp that the member was created at
	JoinedAt string `json:"joinedAt"`

	// (default false)
	IsOwner bool `json:"isOwner,omitempty"`
}

type ServerMemberSummary struct {
	User UserSummary `json:"user"`

	// (must have unique items true)
	RoleIds []int `json:"roleIds"`
}

type ServerMemberPermissions struct {
	Permissions []string `json:"permissions"`
}

type ServerMemberBan struct {
	User UserSummary `json:"user"`

	// The reason for the ban as submitted by the banner
	Reason string `json:"reason,omitempty"`

	// The ID of the user who created this server member ban
	CreatedBy string `json:"createdBy"`

	// The ISO 8601 timestamp that the server member ban was created at
	CreatedAt string `json:"createdAt"`
}

type NicknameResponse struct {
	Nickname string `json:"nickname"`
}

type ServerMemberResponse struct {
	Member ServerMember `json:"member"`
}

type MembersService interface {
	UpdateMemberNickname(userId string, nickname string) (*NicknameResponse, error)
	DeleteMemberNickname(userId string) error
	GetServerMember(serverId string, userId string) (*ServerMember, error)
	KickMember(userId string) error
	BanMember(userId string, reason string) (*ServerMemberBan, error)
	IsMemberBanned(userId string) (*ServerMemberBan, error)
	UnbanMember(userId string) error
	GetServerMembers() (*[]ServerMemberSummary, error)
}

type membersEndpoints struct{}

func (e *membersEndpoints) Nickname(serverId, userId string) string {
	return guildedApi + "/servers" + serverId + "/members" + userId + "/nickname"
}

func (e *membersEndpoints) GetMember(serverId, userId string) string {
	return guildedApi + "/servers/" + serverId + "/members/" + userId
}

func (e *membersEndpoints) GetMembers(serverId string) string {
	return guildedApi + "/servers/" + serverId + "/members"
}

func (e *membersEndpoints) Ban(serverId, userId string) string {
	return guildedApi + "/servers/" + serverId + "/bans/" + userId
}

type membersService struct {
	client    *Client
	endpoints membersEndpoints
}

var _ MembersService = &membersService{}

func (service *membersService) UpdateMemberNickname(userId string, nickname string) (*NicknameResponse, error) {
	endpoint := service.endpoints.Nickname(service.client.ServerID, userId)

	body := &NicknameResponse{
		Nickname: nickname,
	}

	resp, err := service.client.PutRequest(endpoint, body)
	if err != nil {
		return nil, err
	}

	var nick NicknameResponse

	err = json.Unmarshal(resp, &nick)
	if err != nil {
		return nil, err
	}

	return &nick, nil
}

func (service *membersService) DeleteMemberNickname(userId string) error {
	endpoint := service.endpoints.Nickname(service.client.ServerID, userId)

	_, err := service.client.DeleteRequest(endpoint)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to delete member nickname. Error: %s", err.Error()))
	}

	return nil
}

func (service *membersService) GetServerMember(serverId string, userId string) (*ServerMember, error) {
	endpoint := service.endpoints.GetMember(serverId, userId)

	var member ServerMemberResponse
	err := service.client.GetRequestV2(endpoint, &member)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New(fmt.Sprintf("Failed to get member. Error: %s", err.Error()))
	}

	return &member.Member, nil
}

func (service *membersService) KickMember(userId string) error {
	endpoint := service.endpoints.GetMember(service.client.ServerID, userId)

	_, err := service.client.DeleteRequest(endpoint)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to kick member. Error: %s", err.Error()))
	}

	return nil
}

func (service *membersService) BanMember(userId string, reason string) (*ServerMemberBan, error) {
	endpoint := service.endpoints.Ban(service.client.ServerID, userId)

	// No need to build a struct here
	body := map[string]string{
		"reason": reason,
	}

	var ban ServerMemberBan
	err := service.client.PostRequestV2(endpoint, body, &ban)
	if err != nil {
		return nil, err
	}

	return &ban, nil
}

func (service *membersService) IsMemberBanned(userId string) (*ServerMemberBan, error) {
	// Do we want to use the serverID from the config, or manually input it?
	endpoint := service.endpoints.Ban(service.client.ServerID, userId)

	var ban ServerMemberBan
	err := service.client.GetRequestV2(endpoint, &ban)
	if err != nil {
		return nil, err
	}

	return &ban, nil
}

func (service *membersService) UnbanMember(userId string) error {
	endpoint := service.endpoints.Ban(service.client.ServerID, userId)

	_, err := service.client.DeleteRequest(endpoint)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	return nil
}

func (service *membersService) GetServerMembers() (*[]ServerMemberSummary, error) {
	endpoint := service.endpoints.GetMembers(service.client.ServerID)

	var members []ServerMemberSummary
	err := service.client.GetRequestV2(endpoint, &members)
	if err != nil {
		return nil, err
	}

	return &members, nil
}
