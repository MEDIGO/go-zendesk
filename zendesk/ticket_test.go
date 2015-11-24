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

	s.Mux.HandleFunc("/api/v2/tickets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{"ticket": {"id": 35436, "subject": "My printer is on fire!"}}`)
	})

	ticket, err := s.Cl.Tickets.Get(35436)

	assert.NoError(t, err)
	assert.Equal(t, *ticket.Id, int64(35436))
	assert.Equal(t, *ticket.Subject, "My printer is on fire!")
}
