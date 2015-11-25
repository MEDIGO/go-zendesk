package zendesk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserSeviceGet(t *testing.T) {
	s := NewTestSuite()
	defer s.Teardown()

	s.Mux.HandleFunc("/api/v2/users/35436.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)

		fmt.Fprint(w, `{"user": {"id": 35436, "name": "Johnny Agent"}}`)
	})

	found, err := s.Client.Users.Get(35436)
	expected := &User{Id: Int(35436), Name: String("Johnny Agent")}

	assert.NoError(t, err)
	assert.Equal(t, found, expected)
}

func TestUserServiceCreate(t *testing.T) {
	s := NewTestSuite()
	defer s.Teardown()

	input := &User{Name: String("Roger Wilco"), Email: String("roge@example.org")}

	s.Mux.HandleFunc("/api/v2/users.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)

		received := &UserBody{&User{}}
		json.NewDecoder(r.Body).Decode(received)

		assert.Equal(t, input, received.User)

		fmt.Fprint(w, `{"user": {"id": 9873843, "name": "Roger Wilco", "email": "roge@example.org"}}`)
	})

	found, err := s.Client.Users.Create(input)
	expected := &User{Id: Int(9873843), Name: String("Roger Wilco"), Email: String("roge@example.org")}

	assert.NoError(t, err)
	assert.Equal(t, expected, found)
}
