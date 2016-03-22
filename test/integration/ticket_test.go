package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/medigo/go-zendesk/zendesk"
)

func TestTicketCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	user := &zendesk.User{
		Name:  zendesk.String("test-" + randstr(7)),
		Email: zendesk.String("test-" + randstr(7) + "@example.com"),
	}

	user, err = client.CreateUser(user)
	assert.NoError(t, err)

	ticket := &zendesk.Ticket{
		Subject:     zendesk.String("test-" + randstr(7)),
		Description: zendesk.String("test-" + randstr(7)),
		RequesterID: user.ID,
		Tags:        &[]string{"test"},
	}

	created, err := client.CreateTicket(ticket)
	assert.NoError(t, err)
	assert.NotNil(t, created.ID)
	assert.Len(t, *ticket.Tags, 1)

	found, err := client.GetTicket(*created.ID)
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
