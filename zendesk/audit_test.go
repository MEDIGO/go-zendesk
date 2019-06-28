package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListTicketAudits(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	user := randUser(t, client)
	defer client.DeleteUser(*user.ID)

	ticket := randTicket(t, client, user)
	defer client.DeleteTicket(*ticket.ID)

	// assert that a newly created ticket has an audit
	listed, err := client.ListTicketAudits(*ticket.ID, nil)
	require.NoError(t, err)
	require.Len(t, listed.Audits, 1)
}
