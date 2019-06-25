package zendesk

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
