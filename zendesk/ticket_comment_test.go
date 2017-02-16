package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTicketCommentCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	user := randUser(t, client)

	ticket := randTicket(t, client, user)

	// assert that a newly created ticket has a comment
	listed, err := client.ListTicketComments(*ticket.ID)
	require.NoError(t, err)
	require.Len(t, listed, 1)

	// assert that we can add a comment to a ticket
	in := Ticket{
		Comment: &TicketComment{
			Body: String("The smoke is very colorful."),
		},
	}

	ticket, err = client.UpdateTicket(*ticket.ID, &in)
	require.NoError(t, err)

	// assert that we can list the newly created comment
	listed, err = client.ListTicketComments(*ticket.ID)
	require.NoError(t, err)
	require.Len(t, listed, 2)
}
