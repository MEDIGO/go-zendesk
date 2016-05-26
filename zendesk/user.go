package zendesk

import (
	"fmt"
	"time"
)

// User represents a Zendesk user.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/users#content
type User struct {
	ID                  *int64                 `json:"id,omitempty"`
	URL                 *string                `json:"url,omitempty"`
	Name                *string                `json:"name,omitempty"`
	ExternalID          *string                `json:"external_id,omitempty"`
	Alias               *string                `json:"alias,omitempty"`
	CreatedAt           *time.Time             `json:"created_at,omitempty"`
	UpdatedAt           *time.Time             `json:"updated_at,omitempty"`
	Active              *bool                  `json:"active,omitempty"`
	Verified            *bool                  `json:"verified,omitempty"`
	Shared              *bool                  `json:"shared,omitempty"`
	SharedAgent         *bool                  `json:"shared_agent,omitempty"`
	Locale              *string                `json:"locale,omitempty"`
	LocaleID            *int64                 `json:"locale_id,omitempty"`
	TimeZone            *string                `json:"time_zone,omitempty"`
	LastLoginAt         *time.Time             `json:"last_login_at,omitempty"`
	Email               *string                `json:"email,omitempty"`
	Phone               *string                `json:"phone,omitempty"`
	Signature           *string                `json:"signature,omitempty"`
	Details             *string                `json:"details,omitempty"`
	Notes               *string                `json:"notes,omitempty"`
	OrganizationID      *int64                 `json:"organization_id,omitempty"`
	Role                *string                `json:"role,omitempty"`
	CustomerRoleID      *int64                 `json:"custom_role_id,omitempty"`
	Moderator           *bool                  `json:"moderator,omitempty"`
	TicketRestriction   *string                `json:"ticket_restriction,omitempty"`
	OnlyPrivateComments *bool                  `json:"only_private_comments,omitempty"`
	Tags                []string               `json:"tags,omitempty"`
	RestrictedAgent     *bool                  `json:"restricted_agent,omitempty"`
	Suspended           *bool                  `json:"suspended,omitempty"`
	UserFields          map[string]interface{} `json:"user_fields,omitempty"`
}

// ShowUser fetches a user by its ID.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/users#show-user
func (c *client) ShowUser(id int64) (*User, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/users/%d.json", id), out)
	return out.User, err
}

// CreateUser creates a user.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/users#create-user
func (c *client) CreateUser(user *User) (*User, error) {
	in := &APIPayload{User: user}
	out := new(APIPayload)
	err := c.post("/api/v2/users.json", in, out)
	return out.User, err
}

// UpdateUser updates a user.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/users#update-user
func (c *client) UpdateUser(id int64, user *User) (*User, error) {
	in := &APIPayload{User: user}
	out := new(APIPayload)
	err := c.put(fmt.Sprintf("/api/v2/users/%d.json", id), in, out)
	return out.User, err
}

// SearchUsers searches users by name or email address.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/users#search-users
func (c *client) SearchUsers(query string) ([]User, error) {
	out := new(APIPayload)
	err := c.get("/api/v2/users/search.json?query="+query, out)
	return out.Users, err
}
