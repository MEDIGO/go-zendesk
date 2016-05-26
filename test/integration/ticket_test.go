package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

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

func TestUpdateManyTickets(t *testing.T) {
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

	updated, err := client.UpdateManyTickets(updates)
	assert.NoError(t, err)

	for _, ticket := range updated {
		assert.Equal(t, "solved", *ticket.Status)
	}
}
