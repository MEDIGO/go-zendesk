package zendesk

import (
	"fmt"
	"time"
)

type Ticket struct {
	Id              *int       `json:"id,omitempty"`
	URL             *string    `json:"url,omitempty"`
	ExternalId      *string    `json:"external_id,omitempty"`
	Type            *string    `json:"type,omitempty"`
	Subject         *string    `json:"subject,omitempty"`
	RawSubject      *string    `json:"raw_subject,omitempty"`
	Description     *string    `json:"description,omitempty"`
	Priority        *string    `json:"priority,omitempty"`
	Status          *string    `json:"status,omitempty"`
	Recipient       *string    `json:"recipient,omitempty"`
	RequesterId     *int       `json:"requester_id,omitempty"`
	SubmitterId     *int       `json:"submitter_id,omitempty"`
	AssigneeId      *int       `json:"assignee_id,omitempty"`
	OrganizationId  *int       `json:"organization_id,omitempty"`
	GroupId         *int       `json:"group_id,omitempty"`
	CollaboratorIds *[]int     `json:"collaborator_ids,omitempty"`
	ForumTopicId    *int       `json:"forum_topic_id,omitempty"`
	ProblemId       *int       `json:"problem_id,omitempty"`
	HasIncidents    *bool      `json:"has_incidents,omitempty"`
	DueAt           *time.Time `json:"due_at,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}

type TicketBody struct {
	Ticket *Ticket `json:"ticket"`
}

type TicketService struct {
	client *Client
}

func NewTicketService(client *Client) *TicketService {
	return &TicketService{client}
}

func (s *TicketService) Get(id int) (*Ticket, error) {
	res := TicketBody{}
	err := s.client.Get(fmt.Sprintf("/api/v2/tickets/%d.json", id), &res)
	return res.Ticket, err
}

func (s *TicketService) Create(ticket *Ticket) (*Ticket, error) {
	req := TicketBody{ticket}
	res := TicketBody{}
	err := s.client.Post("/api/v2/tickets.json", &req, &res)
	return res.Ticket, err
}
