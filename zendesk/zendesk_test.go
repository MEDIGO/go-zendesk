package zendesk

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	client Client
	server *httptest.Server
	mux    *http.ServeMux
}

func (s *TestSuite) SetupTest() {
	s.mux = http.NewServeMux()
	s.server = httptest.NewServer(s.mux)
	url, err := url.Parse(s.server.URL)
	require.NoError(s.T(), err)

	s.client = &client{
		baseURL:   url,
		userAgent: "test",
		username:  "test",
		password:  "test",
	}
}

func (s *TestSuite) TearDownTest() {
	s.server.Close()
}
