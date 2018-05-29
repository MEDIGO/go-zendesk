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

type RedactedString struct {
	Text *string `json:"text"`
}

func (c *client) ListTicketComments(id int64) ([]TicketComment, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/tickets/%d/comments.json", id), out)
	return out.Comments, err
}

// Redact Comment String removes a string in the comment text
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/ticket_comments#redact-string-in-comment
func (c *client) RedactCommentString(id, ticketID int64, text string) (*TicketComment, error) {
	in := &RedactedString{Text: &text}
	out := new(APIPayload)
	err := c.put(
		fmt.Sprintf("/api/v2/tickets/%d/comments/%d/redact.json", ticketID, id),
		in,
		out)

	return out.Comment, err
}
