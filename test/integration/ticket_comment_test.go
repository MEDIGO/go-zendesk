package integration

import (
	"testing"

	"github.com/MEDIGO/go-zendesk/zendesk"
	"github.com/stretchr/testify/assert"
)

func TestTicketCommentCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	user, err := RandUser(client)
	assert.NoError(t, err)

	ticket, err := RandTicket(client, user)
	assert.NoError(t, err)

	// assert that a newly created ticket has a comment
	listed, err := client.ListTicketComments(*ticket.ID)
	assert.NoError(t, err)
	assert.Len(t, listed, 1)

	// assert that we can add a comment to a ticket
	in := zendesk.Ticket{
		Comment: &zendesk.TicketComment{
			Body: zendesk.String("The smoke is very colorful."),
		},
	}

	ticket, err = client.UpdateTicket(*ticket.ID, &in)
	assert.NoError(t, err)

	// assert that we can list the newly created comment
	listed, err = client.ListTicketComments(*ticket.ID)
	assert.NoError(t, err)
	assert.Len(t, listed, 2)
}
