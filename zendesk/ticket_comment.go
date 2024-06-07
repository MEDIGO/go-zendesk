package zendesk

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
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
	Via         *Via         `json:"via,omitempty"`
}

type RedactedString struct {
	Text *string `json:"text"`
}

func (c *client) ListTicketComments(id int64) ([]TicketComment, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/tickets/%d/comments.json", id), out)
	return out.Comments, err
}

func (c *client) ListTicketCommentsFull(id int64, options *ListOptions, sideloads ...SideLoad) (*ListResponse, error) {
	params, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	sideLoads := &SideLoadOptions{}
	for _, opt := range sideloads {
		opt(sideLoads)
	}
	if len(sideLoads.Include) > 0 {
		params.Set("include", strings.Join(sideLoads.Include, ","))
	}
	out := new(APIPayload)
	err = c.get(fmt.Sprintf("/api/v2/tickets/%d/comments.json?%s", id, params.Encode()), out)

	return &ListResponse{
		Comments:     out.Comments,
		Users:        out.Users,
		Groups:       out.Groups,
		NextPage:     out.NextPage,
		PreviousPage: out.PreviousPage,
		Count:        out.Count,
	}, err
}

// Redact Comment String removes a string in the comment text
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/ticket_comments#redact-string-in-comment
func (c *client) RedactCommentString(id, ticketID int64, text string) (*TicketComment, error) {
	in := &RedactedString{Text: &text}
	out := new(APIPayload)
	err := c.put(
		fmt.Sprintf("/api/v2/tickets/%d/comments/%d/redact", ticketID, id),
		in,
		out)

	return out.Comment, err
}
