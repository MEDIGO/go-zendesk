package zendesk

import "time"

// Group represents a Zendesk group.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/groups
type Group struct {
	ID        *int64     `json:"id,omitempty"`
	URL       *string    `json:"url,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Deleted   *bool      `json:"deleted,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
