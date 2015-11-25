package zendesk

import (
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
