package zendesk

import (
    "time"
)

// OrganizationMembership represents a Zendesk association between an org and a user.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organization_memberships
type OrganizationMembership struct {
    ID                  *int64      `json:"id,omitempty"`
    URL                 *string     `json:"url,omitempty"`
    UserID              *int64      `json:"user_id,omitempty"`
    OrganizationID      *int64      `json:"organization_id,omitempty"`
    Default             *bool       `json:"default,omitempty"`
    CreatedAt           *time.Time  `json:"created_at,omitempty"`
    UpdatedAt           *time.Time  `json:"updated_at,omitempty"`
}

// CreateOrganizationMembership creates an organization membership.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organization_memberships#create-membership
func (c *client) CreateOrganizationMembership(orgMembership *OrganizationMembership) (*OrganizationMembership, error) {
    in := &APIPayload{OrganizationMembership: orgMembership}
    out := new(APIPayload)
    err := c.post("/api/v2/organization_memberships.json", in, out)
    return out.OrganizationMembership, err
}