package zendesk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OrganizationServiceSuite struct {
	TestSuite
}

func (s *OrganizationServiceSuite) TestGet() {
	s.mux.HandleFunc("/api/v2/organizations/35436.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)
		fmt.Fprint(w, `{"organization": {"id": 35436, "name": "One Organization"}}`)
	})

	found, err := s.client.GetOrganization(35436)
	expected := &Organization{ID: Int(35436), Name: String("One Organization")}

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), found, expected)
}

func (s *OrganizationServiceSuite) TestCreate() {
	input := &Organization{Name: String("One Organization")}

	s.mux.HandleFunc("/api/v2/organizations.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "POST", r.Method)

		received := new(APIPayload)
		json.NewDecoder(r.Body).Decode(received)

		assert.Equal(s.T(), input, received.Organization)

		fmt.Fprint(w, `{"organization": {"id": 35436, "name": "One Organization"}}`)
	})

	found, err := s.client.CreateOrganization(input)
	expected := &Organization{ID: Int(35436), Name: String("One Organization")}

	assert.NoError(s.T(), err, expected)
	assert.Equal(s.T(), found, expected)
}

func TestOrganizationServiceSuite(t *testing.T) {
	suite.Run(t, new(OrganizationServiceSuite))
}
