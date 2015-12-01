package zendesk

import (
	"fmt"
	"time"
)

type Ticket struct {
	Id              *int64     `json:"id,omitempty"`
	URL             *string    `json:"url,omitempty"`
	ExternalId      *string    `json:"external_id,omitempty"`
	Type            *string    `json:"type,omitempty"`
	Subject         *string    `json:"subject,omitempty"`
	RawSubject      *string    `json:"raw_subject,omitempty"`
	Description     *string    `json:"description,omitempty"`
	Priority        *string    `json:"priority,omitempty"`
	Status          *string    `json:"status,omitempty"`
	Recipient       *string    `json:"recipient,omitempty"`
	RequesterId     *int64     `json:"requester_id,omitempty"`
	SubmitterId     *int64     `json:"submitter_id,omitempty"`
	AssigneeId      *int64     `json:"assignee_id,omitempty"`
	OrganizationId  *int64     `json:"organization_id,omitempty"`
	GroupId         *int64     `json:"group_id,omitempty"`
	CollaboratorIds *[]int64   `json:"collaborator_ids,omitempty"`
	ForumTopicId    *int64     `json:"forum_topic_id,omitempty"`
	ProblemId       *int64     `json:"problem_id,omitempty"`
	HasIncidents    *bool      `json:"has_incidents,omitempty"`
	DueAt           *time.Time `json:"due_at,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}

type TicketService struct {
	client *Client
}

func NewTicketService(client *Client) *TicketService {
	return &TicketService{client}
}

func (s *TicketService) Get(id int64) (*Ticket, error) {
	out := new(APIPayload)
	err := s.client.Get(fmt.Sprintf("/api/v2/tickets/%d.json", id), out)
	return out.Ticket, err
}

func (s *TicketService) Create(ticket *Ticket) (*Ticket, error) {
	in := &APIPayload{Ticket: ticket}
	out := new(APIPayload)
	err := s.client.Post("/api/v2/tickets.json", in, out)
	return out.Ticket, err
}
