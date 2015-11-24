package zendesk

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

type TestSuite struct {
	Cl  *Client
	Srv *httptest.Server
	Mux *http.ServeMux
}

func NewTestSuite() *TestSuite {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client, _ := NewClient("company", "username", "password")
	url, _ := url.Parse(server.URL)

	client.BaseURL = url

	return &TestSuite{
		Cl:  client,
		Srv: server,
		Mux: mux,
	}
}

func (s *TestSuite) Teardown() {
	s.Srv.Close()
}
