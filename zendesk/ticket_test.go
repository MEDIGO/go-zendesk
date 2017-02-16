package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTicketCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	user := randUser(t, client)

	ticket := &Ticket{
		Subject:     String("My printer is on fire!"),
		Description: String("The smoke is very colorful."),
		RequesterID: user.ID,
		Tags:        []string{"test"},
	}

	created, err := client.CreateTicket(ticket)
	require.NoError(t, err)
	require.NotNil(t, created.ID)
	require.Len(t, ticket.Tags, 1)

	found, err := client.ShowTicket(*created.ID)
	require.NoError(t, err)
	require.Equal(t, created.Subject, found.Subject)
	require.Equal(t, created.RequesterID, found.RequesterID)
	require.Equal(t, created.Tags, found.Tags)

	input := Ticket{
		Status: String("solved"),
	}

	updated, err := client.UpdateTicket(*created.ID, &input)
	require.NoError(t, err)
	require.Equal(t, input.Status, updated.Status)

	requested, err := client.ListRequestedTickets(*user.ID)
	require.NoError(t, err)
	require.Len(t, requested, 1)
	require.Equal(t, created.ID, requested[0].ID)
}

func TestBatchUpdateManyTickets(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	user := randUser(t, client)

	one := randTicket(t, client, user)
	two := randTicket(t, client, user)

	updates := []Ticket{
		{ID: one.ID, Status: String("solved")},
		{ID: two.ID, Status: String("solved")},
	}

	err = client.BatchUpdateManyTickets(updates)
	require.NoError(t, err)
}

func TestBulkUpdateManyTickets(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	user := randUser(t, client)

	one := randTicket(t, client, user)
	require.True(t, contains(one.Tags, "test"))

	two := randTicket(t, client, user)
	require.True(t, contains(two.Tags, "test"))

	err = client.BulkUpdateManyTickets([]int64{*one.ID, *two.ID}, &Ticket{
		AdditionalTags: []string{"a_new_tag"},
		RemoveTags:     []string{"test"},
	})
	require.NoError(t, err)
}

func TestListTicketIncidents(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	user := randUser(t, client)
	ticket := randTicket(t, client, user)

	incident1 := &Ticket{
		Description: String("Fire alarm trigger"),
		RequesterID: user.ID,
		Type:        String("incident"),
		ProblemID:   ticket.ID,
	}

	incident2 := &Ticket{
		Description: String("Building evacuation"),
		RequesterID: user.ID,
		Type:        String("incident"),
		ProblemID:   ticket.ID,
	}

	_, err = client.CreateTicket(incident1)
	require.NoError(t, err)
	_, err = client.CreateTicket(incident2)
	require.NoError(t, err)

	incidents, err := client.ListTicketIncidents(*ticket.ID)

	require.NoError(t, err)
	require.Len(t, incidents, 2)
}

func contains(l []string, s string) bool {
	for _, e := range l {
		if e == s {
			return true
		}
	}
	return false
}
