package zendesk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

// Client describes a client for the Zendesk Core API.
type Client interface {
	BatchUpdateManyTickets([]Ticket) error
	BulkUpdateManyTickets([]int64, *Ticket) error
	CreateOrganization(*Organization) (*Organization, error)
	CreateOrganizationMembership(*OrganizationMembership) (*OrganizationMembership, error)
	CreateOrUpdateUser(*User) (*User, error)
	CreateTicket(*Ticket) (*Ticket, error)
	CreateUser(*User) (*User, error)
	DeleteOrganization(int64) error
	DeleteTicket(int64) error
	DeleteUser(int64) (*User, error)
	DeleteOrganizationMembershipByID(int64) error
	ListLocales() ([]Locale, error)
	ListOrganizationMembershipsByUserID(id int64) ([]OrganizationMembership, error)
	ListOrganizations(*ListOptions) ([]Organization, error)
	ListOrganizationUsers(int64, *ListUsersOptions) ([]User, error)
	ListRequestedTickets(int64) ([]Ticket, error)
	ListTicketComments(int64) ([]TicketComment, error)
	ListTicketFields() ([]TicketField, error)
	ListTicketIncidents(int64) ([]Ticket, error)
	SearchUsers(string) ([]User, error)
	ShowLocale(int64) (*Locale, error)
	ShowLocaleByCode(string) (*Locale, error)
	ShowManyUsers([]int64) ([]User, error)
	ShowOrganization(int64) (*Organization, error)
	ShowTicket(int64) (*Ticket, error)
	ShowUser(int64) (*User, error)
	UpdateOrganization(int64, *Organization) (*Organization, error)
	UpdateTicket(int64, *Ticket) (*Ticket, error)
	UpdateUser(int64, *User) (*User, error)
	UploadFile(string, *string, io.Reader) (*Upload, error)
}

type client struct {
	username string
	password string

	client    *http.Client
	baseURL   *url.URL
	userAgent string
}

// NewEnvClient creates a new Client configured via environment variables.
func NewEnvClient() (Client, error) {
	domain := os.Getenv("ZENDESK_DOMAIN")
	if domain == "" {
		return nil, errors.New("ZENDESK_DOMAIN not found")
	}

	username := os.Getenv("ZENDESK_USERNAME")
	if username == "" {
		return nil, errors.New("ZENDESK_USERNAME not found")
	}

	password := os.Getenv("ZENDESK_PASSWORD")
	if password == "" {
		return nil, errors.New("ZENDESK_PASSWORD not found")
	}

	return NewClient(domain, username, password)
}

// NewClient creates a new Client.
func NewClient(domain, username, password string) (Client, error) {
	baseURL, err := url.Parse(fmt.Sprintf("https://%s.zendesk.com", domain))
	if err != nil {
		return nil, err
	}

	return &client{
		baseURL:   baseURL,
		userAgent: "Go-Zendesk",
		username:  username,
		password:  password,
	}, err
}

func (c *client) request(method, endpoint string, headers map[string]string, body io.Reader) (*http.Response, error) {
	rel, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	url := c.baseURL.ResolveReference(rel)
	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("User-Agent", c.userAgent)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return http.DefaultClient.Do(req)
}

func (c *client) do(method, endpoint string, in, out interface{}) error {
	payload, err := marshall(in)
	if err != nil {
		return err
	}

	headers := map[string]string{}
	if in != nil {
		headers["Content-Type"] = "application/json"
	}

	res, err := c.request(method, endpoint, headers, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Retry the request if the retry after header is present. This can happen when we are
	// being rate limited or we failed with a retriable error.
	if res.Header.Get("Retry-After") != "" {
		after, err := strconv.ParseInt(res.Header.Get("Retry-After"), 10, 64)
		if err != nil || after == 0 {
			return unmarshall(res, out)
		}

		time.Sleep(time.Duration(after) * time.Second)

		res, err = c.request(method, endpoint, headers, bytes.NewReader(payload))
		if err != nil {
			return err
		}
		defer res.Body.Close()
	}

	return unmarshall(res, out)
}

func (c *client) get(endpoint string, out interface{}) error {
	return c.do("GET", endpoint, nil, out)
}

func (c *client) post(endpoint string, in, out interface{}) error {
	return c.do("POST", endpoint, in, out)
}

func (c *client) put(endpoint string, in, out interface{}) error {
	return c.do("PUT", endpoint, in, out)
}

func (c *client) delete(endpoint string, out interface{}) error {
	return c.do("DELETE", endpoint, nil, out)
}

func marshall(in interface{}) ([]byte, error) {
	if in == nil {
		return nil, nil
	}

	return json.Marshal(in)
}

func unmarshall(res *http.Response, out interface{}) error {
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		apierr := new(APIError)
		apierr.Response = res
		if err := json.NewDecoder(res.Body).Decode(apierr); err != nil {
			apierr.Type = String("Unknown")
			apierr.Description = String("Oops! Something went wrong when parsing the error response.")
		}
		return apierr
	}

	if out != nil {
		return json.NewDecoder(res.Body).Decode(out)
	}

	return nil
}

// APIPayload represents the payload of an API call.
type APIPayload struct {
	Attachment              *Attachment              `json:"attachment"`
	Attachments             []Attachment             `json:"attachments"`
	Comment                 *TicketComment           `json:"comment,omitempty"`
	Comments                []TicketComment          `json:"comments,omitempty"`
	Locale                  *Locale                  `json:"locale,omitempty"`
	Locales                 []Locale                 `json:"locales,omitempty"`
	Organization            *Organization            `json:"organization,omitempty"`
	Organizations           []Organization           `json:"organizations,omitempty"`
	Ticket                  *Ticket                  `json:"ticket,omitempty"`
	Tickets                 []Ticket                 `json:"tickets,omitempty"`
	Upload                  *Upload                  `json:"upload,omitempty"`
	User                    *User                    `json:"user,omitempty"`
	Users                   []User                   `json:"users,omitempty"`
	TicketField             *TicketField             `json:"ticket_field,omitempty"`
	TicketFields            []TicketField            `json:"ticket_fields,omitempty"`
	OrganizationMembership  *OrganizationMembership  `json:"organization_membership,omitempty"`
	OrganizationMemberships []OrganizationMembership `json:"organization_memberships,omitempty"`
}

// APIError represents an error response returnted by the API.
type APIError struct {
	Response *http.Response

	Type        *string                       `json:"error,omitmepty"`
	Description *string                       `json:"description,omitempty"`
	Details     *map[string][]*APIErrorDetail `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	msg := fmt.Sprintf("%v %v: %d", e.Response.Request.Method, e.Response.Request.URL, e.Response.StatusCode)

	if e.Type != nil {
		msg = fmt.Sprintf("%s %v", msg, *e.Type)
	}

	if e.Description != nil {
		msg = fmt.Sprintf("%s: %v", msg, *e.Description)
	}

	if e.Details != nil {
		msg = fmt.Sprintf("%s: %+v", msg, *e.Details)
	}

	return msg
}

// APIErrorDetail represents a detail about an APIError.
type APIErrorDetail struct {
	Type        *string `json:"error,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (e *APIErrorDetail) Error() string {
	msg := ""

	if e.Type != nil {
		msg = *e.Type + ": "
	}

	if e.Description != nil {
		msg += *e.Description
	}

	return msg
}

// Bool is a helper function that returns a pointer to the bool value b.
func Bool(b bool) *bool {
	p := b
	return &p
}

// Int is a helper function that returns a pointer to the int value i.
func Int(i int64) *int64 {
	p := i
	return &p
}

// String is a helper function that returns a pointer to the string value s.
func String(s string) *string {
	p := s
	return &p
}

// ListOptions specifies the optional parameters for the list methods that support pagination.
//
// Zendesk Core API doscs: https://developer.zendesk.com/rest_api/docs/core/introduction#pagination
type ListOptions struct {
	// Sets the page of results to retrieve.
	Page int `url:"page,omitempty"`
	// Sets the number of results to include per page.
	PerPage int `url:"per_page,omitempty"`
}
