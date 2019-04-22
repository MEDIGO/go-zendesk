package zendesk

import (
	"fmt"
	"time"
)

// UserIdentity represents a Zendesk user identity.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/user_identities
type UserIdentity struct {
	ID                 *int64     `json:"id,omitempty"`
	URL                *string    `json:"url,omitempty"`
	UserID             *int64     `json:"user_id,omitempty"`
	Type               *string    `json:"type,omitempty"`
	Value              *string    `json:"value,omitempty"`
	Verified           *bool      `json:"verified,omitempty"`
	Primary            *bool      `json:"primary,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
	UndeliverableCount *int64     `json:"undeliverable_count,omitempty"`
	DeliverableState   *string    `json:"deliverable_state,omitempty"`
	SubTypeName        *string    `json:"subtype_name,omitempty"`
}

// ListIdentities lists all user identities.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/user_identities#list-identities
func (c *client) ListIdentities(userID int64) ([]UserIdentity, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/users/%d/identities.json", userID), out)
	return out.Identities, err
}

// ShowIdentity fetches a user identity by its ID and user ID.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/user_identities#show-identity
func (c *client) ShowIdentity(userID, id int64) (*UserIdentity, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/users/%d/identities/%d.json", userID, id), out)
	return out.Identity, err
}

// CreateIdentity creates a user identity.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/user_identities#create-identity
func (c *client) CreateIdentity(userID int64, identity *UserIdentity) (*UserIdentity, error) {
	in := &APIPayload{Identity: identity}
	out := new(APIPayload)
	err := c.post(fmt.Sprintf("/api/v2/users/%d/identities.json", userID), in, out)
	return out.Identity, err
}

// UpdateIdentity updates the value and verified status of a user identity.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/user_identities#update-identity
func (c *client) UpdateIdentity(userID, id int64, identity *UserIdentity) (*UserIdentity, error) {
	in := &APIPayload{Identity: identity}
	out := new(APIPayload)
	err := c.put(fmt.Sprintf("/api/v2/users/%d/identities/%d.json", userID, id), in, out)
	return out.Identity, err
}

// DeleteIdentity deletes a user identity.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/user_identities#delete-identity
func (c *client) DeleteIdentity(userID, id int64) error {
	return c.delete(fmt.Sprintf("/api/v2/users/%d/identities/%d.json", userID, id), nil)
}
