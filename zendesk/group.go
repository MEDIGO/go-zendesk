package zendesk

import (
	"fmt"
	"time"
)

// Group represents a Zendesk Group.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/groups
type Group struct {
	ID        *int64     `json:"id"`
	URL       *string    `json:"url"`
	Deleted   *bool      `json:"deleted"`
	Name      *string    `json:"name"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (c *client) ListGroups() ([]Group, error) {
	out := new(APIPayload)
	err := c.get("/api/v2/groups.json", out)
	return out.Groups, err
}

func (c *client) ShowGroup(id int64) (*Group, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/groups/%d.json", id), out)
	return out.Group, err
}

func (c *client) SearchGroup(name string) (*Group, error) {
	groups, err := c.ListGroups()
	if err != nil {
		return nil, err
	}

	for _, g := range groups {
		if *g.Name == name {
			return &g, err
		}
	}

	return nil, fmt.Errorf("Group with name %s not found", name)
}
