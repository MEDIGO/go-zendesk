package zendesk

import (
	"fmt"
	"net/http"

	"github.com/stretchr/testify/assert"
)

type LocaleServiceSuite struct {
	TestSuite
}

func (s *LocaleServiceSuite) TestList() {
	s.mux.HandleFunc("/api/v2/locales.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)

		fmt.Fprint(w, `{"locales": [{"id": 35436}, {"id": 9873843}]}`)
	})

	found, err := s.client.UserSearch("Gerry")
	expected := []Locale{Locale{ID: Int(35436)}, Locale{ID: Int(9873843)}}

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), found, expected)
}
