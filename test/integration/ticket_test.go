package integration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/MEDIGO/go-zendesk/zendesk"
)

func TestTicketCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	user, err := RandUser(client)
	assert.NoError(t, err)

	ticket := &zendesk.Ticket{
		Subject:     zendesk.String("My printer is on fire!"),
		Description: zendesk.String("The smoke is very colorful."),
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

	input := zendesk.Ticket{
		Status: zendesk.String("solved"),
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

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	user, err := RandUser(client)

	one, err := RandTicket(client, user)
	two, err := RandTicket(client, user)

	updates := []zendesk.Ticket{
		{ID: one.ID, Status: zendesk.String("solved")},
		{ID: two.ID, Status: zendesk.String("solved")},
	}

	err = client.BatchUpdateManyTickets(updates)
	assert.NoError(t, err)

	time.Sleep(10 * time.Second) // wait for the update many job to complete

	one, err = client.ShowTicket(*one.ID)
	assert.NoError(t, err)

	assert.Equal(t, "solved", *one.Status)

	two, err = client.ShowTicket(*two.ID)
	assert.NoError(t, err)

	assert.Equal(t, "solved", *two.Status)
}

func TestBulkUpdateManyTickets(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	user, err := RandUser(client)

	one, err := RandTicket(client, user)
	require.NoError(t, err)
	assert.True(t, contains(one.Tags, "test"))

	two, err := RandTicket(client, user)
	require.NoError(t, err)
	assert.True(t, contains(two.Tags, "test"))

	err = client.BulkUpdateManyTickets([]int64{*one.ID, *two.ID}, &zendesk.Ticket{
		AdditionalTags: []string{"a_new_tag"},
		RemoveTags:     []string{"test"},
	})
	require.NoError(t, err)

	time.Sleep(10 * time.Second) // wait for the update many job to complete

	one, err = client.ShowTicket(*one.ID)
	assert.NoError(t, err)

	assert.True(t, contains(one.Tags, "a_new_tag"))
	assert.False(t, contains(one.Tags, "test"))

	two, err = client.ShowTicket(*two.ID)
	assert.NoError(t, err)

	assert.True(t, contains(two.Tags, "a_new_tag"))
	assert.False(t, contains(two.Tags, "test"))
}

func TestListTicketIncidents(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	user, err := RandUser(client)
	assert.NoError(t, err)

	ticket, err := RandTicket(client, user)
	assert.NoError(t, err)

	incident1 := &zendesk.Ticket{
		Description: zendesk.String("Fire alarm trigger"),
		RequesterID: user.ID,
		Type:        zendesk.String("incident"),
		ProblemID:   ticket.ID,
	}

	incident2 := &zendesk.Ticket{
		Description: zendesk.String("Building evacuation"),
		RequesterID: user.ID,
		Type:        zendesk.String("incident"),
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
