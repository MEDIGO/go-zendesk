package zendesk

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

// TicketAudit represents an audit on a Ticket.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/support/ticket_audits
type TicketAudit struct {
	ID        *int64                 `json:"id,omitempty"`
	TicketID  *int64                 `json:"ticket_id,omitempty"`
	AuthorID  *int64                 `json:"author_id,omitempty"`
	CreatedAt *time.Time             `json:"created_at,omitempty"`
	UpdatedAt *time.Time             `json:"updated_at,omitempty"`
	Events    []interface{}          `json:"events,omitempty"`
	Via       *Via                   `json:"via,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// The via object of a ticket audit or audit event tells you how or why the audit or event was created
//
// Zendesk Via API docs: https://developer.zendesk.com/rest_api/docs/support/ticket_audits#the-via-object
type Via struct {
	Channel *string `json:"channel,omitempty"`
	Source  *Source `json:"source,omitempty"`
}

type Source struct {
	To   *SourceInfo `json:"to,omitempty"`
	From *SourceInfo `json:"from,omitempty"`
	Rel  *string     `json:"rel,omitempty"`
}

type SourceInfo struct {
	Address                          *string       `json:"address,omitempty"`
	Deleted                          *bool         `json:"deleted,omitempty"`
	FacebookID                       *string       `json:"facebook_id,omitempty"`
	FormattedPhone                   *string       `json:"formatted_phone,omitempty"`
	ID                               *int          `json:"id,omitempty"`
	Name                             *string       `json:"name,omitempty"`
	OriginalRecipients               []string      `json:"original_recipients,omitempty"`
	EmailCCs                         []interface{} `json:"email_ccs,omitempty"`
	Phone                            *string       `json:"phone,omitempty"`
	ProfileURL                       *string       `json:"profile_url,omitempty"`
	RegisteredIntegrationServiceName *string       `json:"registered_integration_service_name,omitempty"`
	RevisionID                       *int          `json:"revision_id,omitempty"`
	ServiceInfo                      *string       `json:"service_info,omitempty"`
	Subject                          *string       `json:"subject,omitempty"`
	SupportsChannelback              *bool         `json:"supports_channelback,omitempty"`
	SupportsClickthrough             *bool         `json:"supports_clickthrough,omitempty"`
	TicketID                         *int          `json:"ticket_id,omitempty"`
	Title                            *string       `json:"title,omitempty"`
	TopicID                          *int          `json:"topic_id,omitempty"`
	TopicName                        *string       `json:"topic_name,omitempty"`
	Username                         *string       `json:"username,omitempty"`
}

func (c *client) ListTicketAudits(ticketID int64, options *ListOptions) (*ListResponse, error) {
	params, err := query.Values(options)
	if err != nil {
		return nil, err
	}

	out := new(APIPayload)
	err = c.get(fmt.Sprintf("/api/v2/tickets/%d/audits.json?%s", ticketID, params.Encode()), &out)
	if err != nil {
		return nil, err
	}
	return &ListResponse{
		Audits:       out.Audits,
		NextPage:     out.NextPage,
		PreviousPage: out.PreviousPage,
		Count:        out.Count,
	}, err
}
