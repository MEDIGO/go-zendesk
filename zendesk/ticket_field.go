package zendesk

import (
	"fmt"
	"time"
)

type TicketField struct {
	ID              *int64     `json:"id,omitempty"`
	Type            *string    `json:"string,omitempty"`
	Title           *string    `json:"title,omitempty"`
	Description     *string    `json:"description,omitempty"`
	Position        *int64     `json:"position,omitempty"`
	Active          *bool      `json:"active,omitempty"`
	VisibleInPortal *bool      `json:"visible_in_portal,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}

// ListTicketFields list all availbale custom ticket fields
func (c *client) ListTicketFields() ([]TicketField, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/ticket_fields.json"), out)

	return out.TicketFields, err
}
