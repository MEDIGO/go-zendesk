package zendesk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TicketServiceSuite struct {
	TestSuite
}

func (s *TicketServiceSuite) TestGet() {
	s.mux.HandleFunc("/api/v2/tickets/35436.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)
		fmt.Fprint(w, `{"ticket": {"id": 35436, "subject": "My printer is on fire!"}}`)
	})

	found, err := s.client.Tickets.Get(35436)
	expected := &Ticket{Id: Int(35436), Subject: String("My printer is on fire!")}

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), found, expected)
}

func (s *TicketServiceSuite) TestCreate() {
	input := &Ticket{Subject: String("My printer is on fire!")}

	s.mux.HandleFunc("/api/v2/tickets.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "POST", r.Method)

		received := new(APIPayload)
		json.NewDecoder(r.Body).Decode(received)

		assert.Equal(s.T(), input, received.Ticket)

		fmt.Fprint(w, `{"ticket": {"id": 35436, "subject": "My printer is on fire!"}}`)
	})

	found, err := s.client.Tickets.Create(input)
	expected := &Ticket{Id: Int(35436), Subject: String("My printer is on fire!")}

	assert.NoError(s.T(), err, expected)
	assert.Equal(s.T(), found, expected)
}

func TestTicketServiceSuite(t *testing.T) {
	suite.Run(t, new(TicketServiceSuite))
}
