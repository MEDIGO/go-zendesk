package zendesk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicketServiceGet(t *testing.T) {
	s := NewTestSuite()
	defer s.Teardown()

	s.Mux.HandleFunc("/api/v2/tickets/35436.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{"ticket": {"id": 35436, "subject": "My printer is on fire!"}}`)
	})

	found, err := s.Client.Tickets.Get(35436)
	expected := &Ticket{Id: Int(35436), Subject: String("My printer is on fire!")}

	assert.NoError(t, err)
	assert.Equal(t, found, expected)
}

func TestTicketServiceCreate(t *testing.T) {
	s := NewTestSuite()
	defer s.Teardown()

	input := &Ticket{Subject: String("My printer is on fire!")}

	s.Mux.HandleFunc("/api/v2/tickets.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)

		received := &TicketBody{Ticket: &Ticket{}}
		json.NewDecoder(r.Body).Decode(received)

		assert.Equal(t, input, received.Ticket)

		fmt.Fprint(w, `{"ticket": {"id": 35436, "subject": "My printer is on fire!"}}`)
	})

	found, err := s.Client.Tickets.Create(input)
	expected := &Ticket{Id: Int(35436), Subject: String("My printer is on fire!")}

	assert.NoError(t, err, expected)
	assert.Equal(t, found, expected)
}
