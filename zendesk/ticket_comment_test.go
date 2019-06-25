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
	defer client.DeleteUser(*user.ID)

	ticket := randTicket(t, client, user)
	defer client.DeleteTicket(*ticket.ID)

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

	// assert that we can paginate and include users in listed comments
	listedFull, err := client.ListTicketCommentsFull(*ticket.ID, &ListOptions{PerPage: 10}, IncludeUsers())
	require.NoError(t, err)
	require.Len(t, listedFull.Comments, 2)
	require.Len(t, listedFull.Users, 2)
}

func TestTicketCommentRedaction(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	user := randUser(t, client)
	defer client.DeleteUser(*user.ID)

	ticket := randTicket(t, client, user)
	defer client.DeleteTicket(*ticket.ID)

	// create a comment with a sensitive string
	in := Ticket{
		Comment: &TicketComment{
			Body: String("The credit card number is 4111111111111111"),
		},
	}

	ticket, err = client.UpdateTicket(*ticket.ID, &in)
	require.NoError(t, err)

	listed, err := client.ListTicketComments(*ticket.ID)
	require.NoError(t, err)
	require.Len(t, listed, 2)

	redactedString := "4111111111111111"
	require.Contains(t, *listed[1].Body, redactedString)

	// assert that we can redact a comment with sensitive information in a ticket
	comment, err := client.RedactCommentString(*listed[1].ID, *ticket.ID, redactedString)
	require.NoError(t, err)
	require.NotContains(t, *comment.Body, redactedString)

	// assert that we receive an error if the string is not found
	_, err = client.RedactCommentString(*listed[1].ID, *ticket.ID, "some confidential text")
	require.Error(t, err)
}
