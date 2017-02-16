package zendesk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTicketCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	assert.NoError(t, err)

	user, err := randUser(client)
	assert.NoError(t, err)

	ticket := &Ticket{
		Subject:     String("My printer is on fire!"),
		Description: String("The smoke is very colorful."),
		RequesterID: user.ID,
		Tags:        []string{"test"},
	}

	created, err := client.CreateTicket(ticket)
	assert.NoError(t, err)
	assert.NotNil(t, created.ID)
	assert.Len(t, ticket.Tags, 1)

	found, err := client.ShowTicket(*created.ID)
	assert.NoError(t, err)
	assert.Equal(t, created.Subject, found.Subject)
	assert.Equal(t, created.RequesterID, found.RequesterID)
	assert.Equal(t, created.Tags, found.Tags)

	input := Ticket{
		Status: String("solved"),
	}

	updated, err := client.UpdateTicket(*created.ID, &input)
	assert.NoError(t, err)
	assert.Equal(t, input.Status, updated.Status)

	requested, err := client.ListRequestedTickets(*user.ID)
	assert.NoError(t, err)
	assert.Len(t, requested, 1)
	assert.Equal(t, created.ID, requested[0].ID)
}

func TestBatchUpdateManyTickets(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	assert.NoError(t, err)

	user, err := randUser(client)

	one, err := randTicket(client, user)
	two, err := randTicket(client, user)

	updates := []Ticket{
		{ID: one.ID, Status: String("solved")},
		{ID: two.ID, Status: String("solved")},
	}

	err = client.BatchUpdateManyTickets(updates)
	assert.NoError(t, err)
}

func TestBulkUpdateManyTickets(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	assert.NoError(t, err)

	user, err := randUser(client)

	one, err := randTicket(client, user)
	require.NoError(t, err)
	assert.True(t, contains(one.Tags, "test"))

	two, err := randTicket(client, user)
	require.NoError(t, err)
	assert.True(t, contains(two.Tags, "test"))

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
	assert.NoError(t, err)

	user, err := randUser(client)
	assert.NoError(t, err)

	ticket, err := randTicket(client, user)
	assert.NoError(t, err)

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
	assert.NoError(t, err)
	_, err = client.CreateTicket(incident2)
	assert.NoError(t, err)

	incidents, err := client.ListTicketIncidents(*ticket.ID)

	assert.NoError(t, err)
	assert.Len(t, incidents, 2)
}

func contains(l []string, s string) bool {
	for _, e := range l {
		if e == s {
			return true
		}
	}
	return false
}
