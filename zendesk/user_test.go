package zendesk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceSuite struct {
	TestSuite
}

func (s *UserServiceSuite) TestGet() {
	s.mux.HandleFunc("/api/v2/users/35436.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)

		fmt.Fprint(w, `{"user": {"id": 35436, "name": "Johnny Agent"}}`)
	})

	found, err := s.client.UserGet(35436)
	expected := &User{ID: Int(35436), Name: String("Johnny Agent")}

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), found, expected)
}

func (s *UserServiceSuite) TestCreate() {
	input := &User{Name: String("Roger Wilco"), Email: String("roge@example.org")}

	s.mux.HandleFunc("/api/v2/users.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "POST", r.Method)

		received := new(APIPayload)
		json.NewDecoder(r.Body).Decode(received)

		assert.Equal(s.T(), input, received.User)

		fmt.Fprint(w, `{"user": {"id": 9873843, "name": "Roger Wilco", "email": "roge@example.org"}}`)
	})

	found, err := s.client.UserCreate(input)
	expected := &User{ID: Int(9873843), Name: String("Roger Wilco"), Email: String("roge@example.org")}

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expected, found)
}

func (s *UserServiceSuite) TestUpdate() {
	input := &User{Name: String("Roger Wilco II")}

	s.mux.HandleFunc("/api/v2/users.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "POST", r.Method)

		received := new(APIPayload)
		json.NewDecoder(r.Body).Decode(received)

		assert.Equal(s.T(), input, received.User)

		fmt.Fprint(w, `{"user": {"id": 9873843, "name": "Roger Wilco II"}}`)
	})

	found, err := s.client.UserCreate(input)
	expected := &User{ID: Int(9873843), Name: String("Roger Wilco II")}

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expected, found)
}

func (s *UserServiceSuite) TestSearch() {
	s.mux.HandleFunc("/api/v2/users/search.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)
		assert.Equal(s.T(), "Gerry", r.URL.Query().Get("query"))

		fmt.Fprint(w, `{"users": [{"id": 35436}, {"id": 9873843}]}`)
	})

	found, err := s.client.UserSearch("Gerry")
	expected := []*User{&User{ID: Int(35436)}, &User{ID: Int(9873843)}}

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), found, expected)
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}
