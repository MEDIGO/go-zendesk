package zendesk

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

// Organization represents a Zendesk organization.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organizations
type Organization struct {
	ID                 *int64                 `json:"id,omitempty"`
	URL                *string                `json:"url,omitempty"`
	ExternalID         *string                `json:"external_id,omitempty"`
	Name               *string                `json:"name,omitempty"`
	CreatedAt          *time.Time             `json:"created_at,omitempty"`
	UpdatedAt          *time.Time             `json:"updated_at,omitempty"`
	DomainNames        *[]string              `json:"domain_names,omitempty"`
	Details            *string                `json:"details,omitempty"`
	Notes              *string                `json:"notes,omitempty"`
	GroupID            *int64                 `json:"group_id,omitempty"`
	SharedTickets      *bool                  `json:"shared_tickets,omitempty"`
	SharedComments     *bool                  `json:"shared_comments,omitempty"`
	Tags               *[]string              `json:"tags,omitempty"`
	OrganizationFields map[string]interface{} `json:"organization_fields,omitempty"`
}

// ShowOrganization fetches an organization by its ID.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organizations#show-organization
func (c *client) ShowOrganization(id int64) (*Organization, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/organizations/%d.json", id), out)
	return out.Organization, err
}

// ShowManyOrganizations accepts a comma-separated list of organization ids or external ids.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/support/organizations#show-many-organizations
func (c *client) ShowManyOrganizations(ids []int64) ([]Organization, error) {
	var sids []string
	for _, id := range ids {
		sids = append(sids, strconv.FormatInt(id, 10))
	}

	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/organizations/show_many.json?ids=%s", strings.Join(sids, ",")), out)
	return out.Organizations, err
}

// CreateOrganization creates an organization.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organizations#create-organization
func (c *client) CreateOrganization(org *Organization) (*Organization, error) {
	in := &APIPayload{Organization: org}
	out := new(APIPayload)
	err := c.post("/api/v2/organizations.json", in, out)
	return out.Organization, err
}

// CreateOrUpdateOrganization creates an organization.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organizations#create-or-update-organization
func (c *client) CreateOrUpdateOrganization(org *Organization) (*Organization, error) {
	in := &APIPayload{Organization: org}
	out := new(APIPayload)
	err := c.post("/api/v2/organizations/create_or_update.json", in, out)
	return out.Organization, err
}

// UpdateOrganization updates an organization.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organizations#update-organization
func (c *client) UpdateOrganization(id int64, org *Organization) (*Organization, error) {
	in := &APIPayload{Organization: org}
	out := new(APIPayload)
	err := c.put(fmt.Sprintf("/api/v2/organizations/%d.json", id), in, out)
	return out.Organization, err
}

// ListOrganizations list all organizations.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organizations#list-organizations
func (c *client) ListOrganizations(opts *ListOptions) ([]Organization, error) {
	params, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	out := new(APIPayload)
	err = c.get("/api/v2/organizations.json?"+params.Encode(), out)
	return out.Organizations, err
}

// DeleteOrganization deletes an Organization.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organizations#delete-organization
func (c *client) DeleteOrganization(id int64) error {
	return c.delete(fmt.Sprintf("/api/v2/organizations/%d.json", id), nil)
}

// AutocompleteOrganizations returns an array of organizations whose name starts with the value specified in the name parameter.
// Note: name is case-insensitive
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organizations#autocomplete-organizations
func (c *client) AutocompleteOrganizations(name string) ([]Organization, error) {
	out := new(APIPayload)
	name = url.QueryEscape(name)
	err := c.get("/api/v2/organizations/autocomplete.json?name="+name, out)
	return out.Organizations, err
}

// SearchOrganizationsByExternalID search all organizations by the external_id value of an organization
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organizations#search-organizations-by-external-id
func (c *client) SearchOrganizationsByExternalID(id string) ([]Organization, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/organizations/search.json?external_id=%s", id), out)
	return out.Organizations, err
}
