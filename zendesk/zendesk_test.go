package zendesk

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	client *Client
	server *httptest.Server
	mux    *http.ServeMux
}

func (s *TestSuite) SetupTest() {
	s.mux = http.NewServeMux()
	s.server = httptest.NewServer(s.mux)

	s.client, _ = NewClient("company", "username", "password")
	url, _ := url.Parse(s.server.URL)
	s.client.BaseURL = url
}

func (s *TestSuite) TearDownTest() {
	s.server.Close()
}
