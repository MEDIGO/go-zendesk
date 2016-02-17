package integration

import (
	"testing"

	"github.com/medigo/go-zendesk/zendesk"
	"github.com/stretchr/testify/assert"
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

	user, err = client.UserCreate(user)
	assert.NoError(t, err)

	ticket := &zendesk.Ticket{
		Subject:     zendesk.String("test-" + randstr(7)),
		Description: zendesk.String("test-" + randstr(7)),
		RequesterID: user.ID,
		Tags:        &[]string{"test"},
	}

	created, err := client.TicketCreate(ticket)
	assert.NoError(t, err)
	assert.NotNil(t, created.ID)
	assert.Len(t, *ticket.Tags, 1)

	found, err := client.TicketGet(*created.ID)
	assert.NoError(t, err)
	assert.Equal(t, created.Subject, found.Subject)
	assert.Equal(t, created.RequesterID, found.RequesterID)
	assert.Equal(t, created.Tags, found.Tags)
}
