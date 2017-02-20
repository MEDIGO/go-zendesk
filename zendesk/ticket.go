package zendesk

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Ticket represents a Zendesk Ticket.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/tickets
type Ticket struct {
	ID              *int64         `json:"id,omitempty"`
	URL             *string        `json:"url,omitempty"`
	ExternalID      *string        `json:"external_id,omitempty"`
	Type            *string        `json:"type,omitempty"`
	Subject         *string        `json:"subject,omitempty"`
	RawSubject      *string        `json:"raw_subject,omitempty"`
	Description     *string        `json:"description,omitempty"`
	Comment         *TicketComment `json:"comment,omitempty"`
	Priority        *string        `json:"priority,omitempty"`
	Status          *string        `json:"status,omitempty"`
	Recipient       *string        `json:"recipient,omitempty"`
	RequesterID     *int64         `json:"requester_id,omitempty"`
	SubmitterID     *int64         `json:"submitter_id,omitempty"`
	AssigneeID      *int64         `json:"assignee_id,omitempty"`
	OrganizationID  *int64         `json:"organization_id,omitempty"`
	GroupID         *int64         `json:"group_id,omitempty"`
	CollaboratorIDs []int64        `json:"collaborator_ids,omitempty"`
	ForumTopicID    *int64         `json:"forum_topic_id,omitempty"`
	ProblemID       *int64         `json:"problem_id,omitempty"`
	HasIncidents    *bool          `json:"has_incidents,omitempty"`
	DueAt           *time.Time     `json:"due_at,omitempty"`
	Tags            []string       `json:"tags,omitempty"`
	Via             *Via           `json:"via,omitempty"`
	CreatedAt       *time.Time     `json:"created_at,omitempty"`
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`
	CustomFields    []CustomField  `json:"custom_fields,omitempty"`

	AdditionalTags []string `json:"additional_tags,omitempty"`
	RemoveTags     []string `json:"remove_tags,omitempty"`
}

type CustomField struct {
	ID    *int64      `json:"id"`
	Value interface{} `json:"value"`
}

func (c *client) ShowTicket(id int64) (*Ticket, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/tickets/%d.json", id), out)
	return out.Ticket, err
}

func (c *client) CreateTicket(ticket *Ticket) (*Ticket, error) {
	in := &APIPayload{Ticket: ticket}
	out := new(APIPayload)
	err := c.post("/api/v2/tickets.json", in, out)
	return out.Ticket, err
}

func (c *client) UpdateTicket(id int64, ticket *Ticket) (*Ticket, error) {
	in := &APIPayload{Ticket: ticket}
	out := new(APIPayload)
	err := c.put(fmt.Sprintf("/api/v2/tickets/%d.json", id), in, out)
	return out.Ticket, err
}

func (c *client) BatchUpdateManyTickets(tickets []Ticket) error {
	in := &APIPayload{Tickets: tickets}
	out := new(APIPayload)
	err := c.put("/api/v2/tickets/update_many.json", in, out)
	return err
}

func (c *client) BulkUpdateManyTickets(ids []int64, ticket *Ticket) error {
	parsed := []string{}
	for _, id := range ids {
		parsed = append(parsed, strconv.FormatInt(id, 10))
	}

	in := &APIPayload{Ticket: ticket}
	out := new(APIPayload)
	err := c.put(fmt.Sprintf("/api/v2/tickets/update_many.json?ids=%s", strings.Join(parsed, ",")), in, out)
	return err
}

func (c *client) ListRequestedTickets(userID int64) ([]Ticket, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/users/%d/tickets/requested.json", userID), out)
	return out.Tickets, err
}

// ListTicketIncidents list all incidents related to the problem
func (c *client) ListTicketIncidents(problemID int64) ([]Ticket, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/tickets/%d/incidents.json", problemID), out)

	return out.Tickets, err
}

// DeleteTickets deletes a Ticket.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/tickets#delete-ticket
func (c *client) DeleteTicket(id int64) error {
	return c.delete(fmt.Sprintf("/api/v2/tickets/%d.json", id), nil)
}
