package zendesk

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

// Status returns tickets set to the specified status.
// Possible values include new, open, pending, hold, solved, or closed
type Status string

const (
	StatusNew     Status = "new"
	StatusOpen    Status = "open"
	StatusPending Status = "pending"
	StatusHold    Status = "hold"
	StatusSolved  Status = "solved"
	StatusClosed  Status = "closed"
)

// ResultType returns records of the specified resource type
// Possible values include ticket, user, organization, or group
type ResultType string

const (
	ResultTypeTicket       ResultType = "ticket"
	ResultTypeUser         ResultType = "user"
	ResultTypeOrganization ResultType = "organization"
	ResultTypeGroup        ResultType = "group"
)

// SearchOperator represents supported search operators for zendesk searches
type SearchOperator string

const (
	Equality             SearchOperator = ":"
	LessThan             SearchOperator = "<"
	GreaterThan          SearchOperator = ">"
	LessThanOrEqualTo    SearchOperator = "<="
	GreaterThanOrEqualTo SearchOperator = ">="
)

// QueryOptions to narrow search results
type QueryOptions struct {
	Search []string
}

// SearchString to pass to zendesk
type Filters func(*QueryOptions)

// StatusFilter filters tickets by their status
func StatusFilter(s Status, o SearchOperator) Filters {
	return func(c *QueryOptions) {
		c.Search = append(c.Search, fmt.Sprintf("status%s%s", o, s))
	}
}

// OrganizationFilter filters tickets for the matching organization
func OrganizationFilter(organizationID int) Filters {
	return func(c *QueryOptions) {
		c.Search = append(c.Search, fmt.Sprintf("organization_id:%s", strconv.Itoa(organizationID)))
	}
}

// GroupNameFilter filters results by their group name
func GroupNameFilter(name string) Filters {
	return func(c *QueryOptions) {
		if strings.Contains(name, " ") {
			name = fmt.Sprintf("\"%s\"", name)
		}
		c.Search = append(c.Search, fmt.Sprintf("group:%s", name))
	}
}

// SearchTickets leverages the unified search api to return tickets
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/support/search
func (c *client) SearchTickets(term string, options *ListOptions, filters ...Filters) (*TicketSearchResults, error) {
	params, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	searchOptions := &QueryOptions{}
	for _, opt := range filters {
		opt(searchOptions)
	}
	queryString := fmt.Sprintf("type:%s ", ResultTypeTicket)
	queryString += strings.Join(searchOptions.Search, " ")
	if term != "" {
		queryString = fmt.Sprintf(`%s /"%s/"`, queryString, term)
	}
	params.Set("query", queryString)
	out := new(TicketSearchResults)
	err = c.get(fmt.Sprintf("/api/v2/search.json?%s", params.Encode()), out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchUsersEx leverages the unified search api to return users
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/support/search
func (c *client) SearchUsersEx(term string, options *ListOptions, filters ...Filters) (*UserSearchResults, error) {
	params, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	searchOptions := &QueryOptions{}
	for _, opt := range filters {
		opt(searchOptions)
	}
	queryString := fmt.Sprintf("type:%s ", ResultTypeUser)
	queryString += strings.Join(searchOptions.Search, " ")
	if term != "" {
		queryString = fmt.Sprintf(`%s /"%s/"`, queryString, term)
	}
	params.Set("query", queryString)
	out := new(UserSearchResults)
	err = c.get(fmt.Sprintf("/api/v2/search.json?%s", params.Encode()), out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
