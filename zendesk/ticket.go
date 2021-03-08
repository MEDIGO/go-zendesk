package zendesk

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

// Ticket represents a Zendesk Ticket.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/tickets
type Ticket struct {
	ID                      *int64         `json:"id,omitempty"`
	URL                     *string        `json:"url,omitempty"`
	ExternalID              *string        `json:"external_id,omitempty"`
	Type                    *string        `json:"type,omitempty"`
	Subject                 *string        `json:"subject,omitempty"`
	RawSubject              *string        `json:"raw_subject,omitempty"`
	Description             *string        `json:"description,omitempty"`
	Comment                 *TicketComment `json:"comment,omitempty"`
	CommentCount            *int64         `json:"comment_count,omitempty"`
	Priority                *string        `json:"priority,omitempty"`
	Status                  *string        `json:"status,omitempty"`
	Recipient               *string        `json:"recipient,omitempty"`
	RequesterID             *int64         `json:"requester_id,omitempty"`
	Requester               *Requester     `json:"requester,omitempty"`
	SubmitterID             *int64         `json:"submitter_id,omitempty"`
	AssigneeID              *int64         `json:"assignee_id,omitempty"`
	AssigneeEmail           *string        `json:"assignee_email,omitempty"`
	OrganizationID          *int64         `json:"organization_id,omitempty"`
	GroupID                 *int64         `json:"group_id,omitempty"`
	CollaboratorIDs         []int64        `json:"collaborator_ids,omitempty"`
	Collaborators           []interface{}  `json:"collaborators,omitempty"`
	AdditionalCollaborators []interface{}  `json:"additional_collaborators,omitempty"`
	ForumTopicID            *int64         `json:"forum_topic_id,omitempty"`
	ProblemID               *int64         `json:"problem_id,omitempty"`
	HasIncidents            *bool          `json:"has_incidents,omitempty"`
	DueAt                   *time.Time     `json:"due_at,omitempty"`
	Tags                    []string       `json:"tags,omitempty"`
	Via                     *Via           `json:"via,omitempty"`
	CreatedAt               *time.Time     `json:"created_at,omitempty"`
	UpdatedAt               *time.Time     `json:"updated_at,omitempty"`
	CustomFields            []CustomField  `json:"custom_fields,omitempty"`
	BrandID                 *int64         `json:"brand_id,omitempty"`
	TicketFormID            *int64         `json:"ticket_form_id,omitempty"`
	FollowupSourceID        *int64         `json:"via_followup_source_id,omitempty"`

	AdditionalTags []string `json:"additional_tags,omitempty"`
	RemoveTags     []string `json:"remove_tags,omitempty"`
}

type BulkImportTicket struct {
	AssigneeID          *int64              `json:"assignee_id,omitempty"`
	Comments            []TicketComment     `json:"comments,omitempty"`
	CollaboratorIDs     []int               `json:"collaborator_ids,omitempty"`
	CreatedAt           time.Time           `json:"created_at,omitempty"`
	CustomFields        []CustomField       `json:"custom_fields,omitempty"`
	Description         string              `json:"description,omitempty"`
	DueAt               *time.Time          `json:"due_at,omitempty"`
	ExternalID          string              `json:"external_id,omitempty"`
	FollowerIds         []int               `json:"follower_ids,omitempty"`
	GroupID             *int64              `json:"group_id,omitempty"`
	HasIncidents        bool                `json:"has_incidents,omitempty"`
	ID                  int                 `json:"id,omitempty"`
	OrganizationID      *int64              `json:"organization_id,omitempty"`
	Priority            string              `json:"priority,omitempty"`
	ProblemID           int                 `json:"problem_id,omitempty"`
	RawSubject          string              `json:"raw_subject,omitempty"`
	Recipient           string              `json:"recipient,omitempty"`
	RequesterID         *int64              `json:"requester_id,omitempty"`
	SatisfactionRating  *SatisfactionRating `json:"satisfaction_rating,omitempty"`
	SharingAgreementIds []int               `json:"sharing_agreement_ids,omitempty"`
	Status              string              `json:"status,omitempty"`
	Subject             string              `json:"subject,omitempty"`
	SubmitterID         int                 `json:"submitter_id,omitempty"`
	Tags                []string            `json:"tags,omitempty"`
	Type                string              `json:"type,omitempty"`
	SolvedAt            time.Time           `json:"solve_at,omitempty"`
	UpdatedAt           time.Time           `json:"updated_at,omitempty"`
	URL                 string              `json:"url,omitempty"`
	Via                 Via                 `json:"via,omitempty"`
}

type SatisfactionRating struct {
	Comment string `json:"comment"`
	ID      int    `json:"id"`
	Score   string `json:"score"`
}

type CustomField struct {
	ID    *int64      `json:"id"`
	Value interface{} `json:"value"`
}

type Collaborator struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type Requester struct {
	LocaleID *int    `json:"locale_id"`
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
}

type BulkPayload struct {
	Ticket *BulkImportTicket `json:"ticket,omitempty"`
}

func (c *client) BulkImportTicket(bulk *BulkImportTicket) (*BulkImportTicket, error) {
	in := &BulkPayload{Ticket: bulk}
	out := new(BulkPayload)
	err := c.post("/api/v2/imports/tickets.json", in, out)
	return out.Ticket, err
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
	var parsed []string
	for _, id := range ids {
		parsed = append(parsed, strconv.FormatInt(id, 10))
	}

	in := &APIPayload{Ticket: ticket}
	out := new(APIPayload)
	err := c.put(fmt.Sprintf("/api/v2/tickets/update_many.json?ids=%s", strings.Join(parsed, ",")), in, out)
	return err
}

// ListOrganizationTickets list tickets for an organization
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/tickets#list-tickets
func (c *client) ListOrganizationTickets(organizationID int64, options *ListOptions, sideloads ...SideLoad) (*ListResponse, error) {
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
	err = c.get(fmt.Sprintf("/api/v2/organizations/%d/tickets.json?%s", organizationID, params.Encode()), out)
	if err != nil {
		return nil, err
	}
	return &ListResponse{
		Tickets:      out.Tickets,
		Users:        out.Users,
		Groups:       out.Groups,
		NextPage:     out.NextPage,
		PreviousPage: out.PreviousPage,
		Count:        out.Count,
	}, err
}

// ListExternalIDTickets list tickets by external ID
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/support/tickets#list-tickets-by-external-id
func (c *client) ListExternalIDTickets(externalID string, options *ListOptions, sideloads ...SideLoad) (*ListResponse, error) {
	params, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	params.Set("external_id", externalID)

	sideLoads := &SideLoadOptions{}
	for _, opt := range sideloads {
		opt(sideLoads)
	}
	if len(sideLoads.Include) > 0 {
		params.Set("include", strings.Join(sideLoads.Include, ","))
	}
	out := new(APIPayload)
	err = c.get(fmt.Sprintf("/api/v2/tickets.json?%s", params.Encode()), out)
	if err != nil {
		return nil, err
	}
	return &ListResponse{
		Tickets:      out.Tickets,
		Users:        out.Users,
		Groups:       out.Groups,
		NextPage:     out.NextPage,
		PreviousPage: out.PreviousPage,
		Count:        out.Count,
	}, err
}

// ListRequestedTickets lists tickets that the requesting agent recently viewed in the agent interface,
// not recently created or updated tickets (unless by the agent recently in the agent interface).
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/tickets#list-tickets
func (c *client) ListRequestedTickets(userID int64) ([]Ticket, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/users/%d/tickets/requested.json", userID), out)
	return out.Tickets, err
}

// ListTicketCollaborators lists collaborators by ticket id
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/support/tickets#list-collaborators-for-a-ticket
func (c *client) ListTicketCollaborators(ticketID int64) ([]User, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/tickets/%d/collaborators.json", ticketID), out)
	return out.Users, err
}

// ListTicketFollowers lists followers by ticket id
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/support/tickets#list-followers-for-a-ticket
func (c *client) ListTicketFollowers(ticketID int64) ([]User, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/tickets/%d/followers.json", ticketID), out)
	return out.Users, err
}

// ListTicketEmailCCs lists email CCs by ticket id
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/support/tickets#list-email-ccs-for-a-ticket
func (c *client) ListTicketEmailCCs(ticketID int64) ([]User, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/tickets/%d/email_ccs.json", ticketID), out)
	return out.Users, err
}

// ListTicketIncidents list all incidents related to the problem
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/tickets#listing-ticket-incidents
func (c *client) ListTicketIncidents(problemID int64) ([]Ticket, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/tickets/%d/incidents.json", problemID), out)

	return out.Tickets, err
}

// ListTickets lists all tickets
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/support/tickets#list-tickets
func (c *client) ListTickets(options *ListOptions, sideloads ...SideLoad) (*ListResponse, error) {
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
	err = c.get(fmt.Sprintf("/api/v2/tickets.json?%s", params.Encode()), out)
	if err != nil {
		return nil, err
	}
	return &ListResponse{
		Tickets:      out.Tickets,
		Users:        out.Users,
		Groups:       out.Groups,
		NextPage:     out.NextPage,
		PreviousPage: out.PreviousPage,
		Count:        out.Count,
	}, err
}

// DeleteTickets deletes a Ticket.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/tickets#delete-ticket
func (c *client) DeleteTicket(id int64) error {
	return c.delete(fmt.Sprintf("/api/v2/tickets/%d.json", id), nil)
}

// PermanentlyDeleteTicket purges a ticket with all it's associated data - recordings & attachments
// WARNING: this task is irreversible; GDPR compliant
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/tickets#delete-tickets-permanently
func (c *client) PermanentlyDeleteTicket(id int64) (*JobStatus, error) {
	out := new(APIPayload)
	err := c.delete(fmt.Sprintf("/api/v2/deleted_tickets/%d.json", id), out)
	return out.JobStatus, err
}
