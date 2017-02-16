package zendesk

import (
	"fmt"
	"time"
)

// TicketComment represents a comment on a Ticket.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/ticket_comments
type TicketComment struct {
	ID          *int64       `json:"id,omitempty"`
	Type        *string      `json:"type,omitempty"`
	Body        *string      `json:"body,omitempty"`
	HTMLBody    *string      `json:"html_body,omitempty"`
	Public      *bool        `json:"public,omitempty"`
	AuthorID    *int64       `json:"author_id,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	CreatedAt   *time.Time   `json:"created_at,omitempty"`
	Uploads     []string     `json:"uploads,omitempty"`
}

func (c *client) ListTicketComments(id int64) ([]TicketComment, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/tickets/%d/comments.json", id), out)
	return out.Comments, err
}
