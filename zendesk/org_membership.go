package zendesk

import (
	"fmt"
	"time"
)

// OrganizationMembership represents a Zendesk association between an org and a user.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organization_memberships
type OrganizationMembership struct {
	ID             *int64     `json:"id,omitempty"`
	URL            *string    `json:"url,omitempty"`
	UserID         *int64     `json:"user_id,omitempty"`
	OrganizationID *int64     `json:"organization_id,omitempty"`
	Default        *bool      `json:"default,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
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

// ListOrganizationMembershipsByUserID returns all organization memberships for a specific user
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organization_memberships#list-memberships
func (c *client) ListOrganizationMembershipsByUserID(id int64) ([]OrganizationMembership, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/users/%d/organization_memberships.json", id), out)
	return out.OrganizationMemberships, err
}

// ListOrganizationMembershipsByOrganisationID returns all organization memberships for a specific organisation
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organization_memberships#list-memberships
func (c *client) ListOrganizationMembershipsByOrganisationID(id int64) ([]OrganizationMembership, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/organizations/%d/organization_memberships.json", id), out)
	return out.OrganizationMemberships, err
}

// DeleteOrganizationMembership removes an organization membership
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organization_memberships#delete-membership
func (c *client) DeleteOrganizationMembershipByID(id int64) error {
	return c.delete(fmt.Sprintf("/api/v2/organization_memberships/%d.json", id), nil)
}
