package zendesk

import (
	"fmt"
	"time"
)

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

// ShowGroup fetches a group by its ID.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/groups#show-group
func (c *client) ShowGroup(id int64) (*Group, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/groups/%d.json", id), out)
	return out.Group, err
}

// CreateGroup creates a group.
func (c *client) CreateGroup(group *Group) (*Group, error) {
	in := &APIPayload{Group: group}
	out := new(APIPayload)
	err := c.post("/api/v2/groups.json", in, out)
	return out.Group, err
}

// ListGroups lists all groups.
func (c *client) ListGroups() ([]Group, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/groups.json"), out)

	return out.Groups, err
}

// UpdateGroup updates a group.
func (c *client) UpdateGroup(id int64, group *Group) (*Group, error) {
	in := &APIPayload{Group: group}
	out := new(APIPayload)
	err := c.put(fmt.Sprintf("/api/v2/groups/%d.json", id), in, out)
	return out.Group, err
}

// DeleteGroup deletes a group.
func (c *client) DeleteGroup(id int64) error {
	err := c.delete(fmt.Sprintf("/api/v2/groups/%d.json", id), nil)
	return err
}
