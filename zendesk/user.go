package zendesk

import (
	"fmt"
	"time"
)

type User struct {
	Id                  *int64     `json:"id,omitempty"`
	Url                 *string    `json:"url,omitempty"`
	Name                *string    `json:"name,omitempty"`
	ExternalId          *string    `json:"external_id,omitempty"`
	Alias               *string    `json:"alias,omitempty"`
	CreatedAt           *time.Time `json:"created_at,omitempty"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
	Active              *bool      `json:"active,omitempty"`
	Verified            *bool      `json:"verified,omitempty"`
	Shared              *bool      `json:"shared,omitempty"`
	SharedAgent         *bool      `json:"shared_agent,omitempty"`
	Locale              *string    `json:"locale,omitempty"`
	LocaleId            *int64     `json:"locale_id,omitempty"`
	Time_zone           *string    `json:"time_zone,omitempty"`
	LastLoginAt         *time.Time `json:"last_login_at,omitempty"`
	Email               *string    `json:"email,omitempty"`
	Phone               *string    `json:"phone,omitempty"`
	Signature           *string    `json:"signature,omitempty"`
	Details             *string    `json:"details,omitempty"`
	Notes               *string    `json:"notes,omitempty"`
	OrganizationId      *int64     `json:"organization_id,omitempty"`
	Role                *string    `json:"role,omitempty"`
	CustomerRoleId      *int64     `json:"custom_role_id,omitempty"`
	Moderator           *bool      `json:"moderator,omitempty"`
	TicketRestriction   *string    `json:"ticket_restriction,omitempty"`
	OnlyPrivateComments *bool      `json:"only_private_comments,omitempty"`
	Tags                []string   `json:"tags,omitempty"`
	RestrictedAgent     *bool      `json:"restricted_agent,omitempty"`
	Suspended           *bool      `json:"suspended,omitempty"`
}

type UserBody struct {
	User *User `json:"user"`
}

type UserService struct {
	client *Client
}

func NewUserService(client *Client) *UserService {
	return &UserService{client}
}

func (s *UserService) Get(id int64) (*User, error) {
	res := UserBody{}

	err := s.client.Get(fmt.Sprintf("/api/v2/users/%d.json", id), &res)
	return res.User, err
}

func (s *UserService) Create(user *User) (*User, error) {
	req := UserBody{user}
	res := UserBody{}

	err := s.client.Post("/api/v2/users.json", &req, &res)
	return res.User, err
}
