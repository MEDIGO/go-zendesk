package zendesk

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
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

func (c *client) ShowManyUsers(ids []int64) ([]User, error) {
	sids := []string{}
	for _, id := range ids {
		sids = append(sids, strconv.FormatInt(id, 10))
	}

	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/users/show_many.json?ids=%s", strings.Join(sids, ",")), out)
	return out.Users, err
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

// CreateOrUpdateUser creates or updates a user.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/users#create-or-update-user
func (c *client) CreateOrUpdateUser(user *User) (*User, error) {
	in := &APIPayload{User: user}
	out := new(APIPayload)
	err := c.post("/api/v2/users/create_or_update.json", in, out)
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

// DeleteUser deletes an User.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/users#delete-user
func (c *client) DeleteUser(id int64) (*User, error) {
	out := new(APIPayload)
	err := c.delete(fmt.Sprintf("/api/v2/users/%d.json", id), out)
	return out.User, err
}

// ListUsersOptions specifies the optional parameters for the list users methods.
type ListUsersOptions struct {
	ListOptions

	Role          []string `url:"role"`
	PermissionSet int64    `url:"permision_set"`
}

// ListOrganizationUsers list the users associated to an organization.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/users#list-users
func (c *client) ListOrganizationUsers(id int64, opts *ListUsersOptions) ([]User, error) {
	params, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	out := new(APIPayload)
	err = c.get(fmt.Sprintf("/api/v2/organizations/%d/users.json?%s", id, params.Encode()), out)
	return out.Users, err
}

// SearchUsers searches users by name or email address.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/users#search-users
func (c *client) SearchUsers(query string) ([]User, error) {
	out := new(APIPayload)
	err := c.get("/api/v2/users/search.json?query="+query, out)
	return out.Users, err
}
