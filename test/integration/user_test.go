package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/medigo/go-zendesk/zendesk"
)

func TestUserCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	input := zendesk.User{
		Name:  zendesk.String("test-" + randstr(7)),
		Email: zendesk.String("test-" + randstr(7) + "@example.com"),
	}

	created, err := client.Users.Create(&input)
	assert.NoError(t, err)
	assert.NotNil(t, created.Id)
	assert.Equal(t, *input.Name, *created.Name)
	assert.Equal(t, *input.Email, *created.Email)

	found, err := client.Users.Get(*created.Id)
	assert.NoError(t, err)
	assert.Equal(t, *created.Id, *found.Id)
	assert.Equal(t, *created.Name, *found.Name)
	assert.Equal(t, *created.Email, *found.Email)
}
