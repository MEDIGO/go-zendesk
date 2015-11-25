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

	user, err = client.Users.Create(user)
	assert.NoError(t, err)

	ticket := &zendesk.Ticket{
		Subject:     zendesk.String("test-" + randstr(7)),
		Description: zendesk.String("test-" + randstr(7)),
		RequesterId: user.Id,
	}

	created, err := client.Tickets.Create(ticket)
	assert.NoError(t, err)
	assert.NotNil(t, created.Id)

	found, err := client.Tickets.Get(*created.Id)
	assert.NoError(t, err)
	assert.Equal(t, created.Subject, found.Subject)
	assert.Equal(t, created.RequesterId, found.RequesterId)
}
