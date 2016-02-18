package zendesk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type Client interface {
	LocaleGet(int64) (*Locale, error)
	LocaleGetByCode(string) (*Locale, error)
	LocaleList() ([]Locale, error)
	OrganizationCreate(*Organization) (*Organization, error)
	OrganizationGet(int64) (*Organization, error)
	TicketCreate(*Ticket) (*Ticket, error)
	TicketGet(int64) (*Ticket, error)
	UserCreate(*User) (*User, error)
	UserGet(int64) (*User, error)
	UserSearch(string) ([]User, error)
	UserUpdate(int64, *User) (*User, error)
}

type client struct {
	username string
	password string

	client    *http.Client
	baseURL   *url.URL
	userAgent string
}

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

func (c *client) do(method, endpoint string, in interface{}, out interface{}) error {
	rel, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	url := c.baseURL.ResolveReference(rel)
	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("User-Agent", c.userAgent)

	if in != nil {
		payload, err := json.Marshal(in)
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer(payload)
		req.Body = ioutil.NopCloser(buf)

		req.ContentLength = int64(len(payload))
		req.Header.Set("Content-Length", strconv.Itoa(len(payload)))
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if code := res.StatusCode; 200 <= code && code <= 299 {
		if out != nil {
			return json.NewDecoder(res.Body).Decode(out)
		} else {
			return nil
		}
	}

	apierr := new(APIError)
	apierr.Response = res
	json.NewDecoder(res.Body).Decode(apierr)

	return apierr
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

type APIPayload struct {
	User          *User          `json:"user,omitempty"`
	Users         []User         `json:"users,omitempty"`
	Ticket        *Ticket        `json:"ticket,omitempty"`
	Tickets       []Ticket       `json:"tickets,omitempty"`
	Organization  *Organization  `json:"organization,omitempty"`
	Organizations []Organization `json:"organizations,omitempty"`
	Locale        *Locale        `json:"locale,omitempty"`
	Locales       []Locale       `json:"locales,omitempty"`
}

type APIError struct {
	Response *http.Response

	Type        *string                       `json:"error,omitmepty"`
	Description *string                       `json:"description,omitempty"`
	Details     *map[string][]*APIErrorDetail `json: "details,omitempty"`
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

type APIErrorDetail struct {
	Type        *string `json:"error,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (e *APIErrorDetail) Error() string {
	return fmt.Sprintf("%s: %s", *e.Type, *e.Description)
}

func Bool(b bool) *bool {
	p := b
	return &p
}

func Int(i int64) *int64 {
	p := i
	return &p
}

func String(s string) *string {
	p := s
	return &p
}
