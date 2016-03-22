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

	found, err := s.client.SearchUsers("Gerry")
	expected := []Locale{Locale{ID: Int(35436)}, Locale{ID: Int(9873843)}}

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), found, expected)
}

func (s *LocaleServiceSuite) TestGet() {
	s.mux.HandleFunc("/api/v2/locales/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)

		fmt.Fprint(w, `{"locale": {"id": 1}}`)
	})

	found, err := s.client.GetLocale(1)
	expected := &Locale{ID: Int(1)}

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), found, expected)
}

func (s *LocaleServiceSuite) TestGetByCode() {
	s.mux.HandleFunc("/api/v2/locales/en-US.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)

		fmt.Fprint(w, `{"locale": {"locale": "en-US"}}`)
	})

	found, err := s.client.GetLocaleByCode("en-US")
	expected := &Locale{Locale: String("en-US")}

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), found, expected)
}
