package zendesk

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

type TestSuite struct {
	Client *Client
	Server *httptest.Server
	Mux    *http.ServeMux
}

func NewTestSuite() *TestSuite {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client, _ := NewClient("company", "username", "password")
	url, _ := url.Parse(server.URL)

	client.BaseURL = url

	return &TestSuite{
		Client: client,
		Server: server,
		Mux:    mux,
	}
}

func (s *TestSuite) Teardown() {
	s.Server.Close()
}
