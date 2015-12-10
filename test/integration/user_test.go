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
		UserFields: map[string]interface{}{
			"test": "this is a test",
		},
	}

	created, err := client.UserCreate(&input)
	assert.NoError(t, err)
	assert.NotNil(t, created.Id)
	assert.Equal(t, *input.Name, *created.Name)
	assert.Equal(t, *input.Email, *created.Email)

	found, err := client.UserGet(*created.Id)
	assert.NoError(t, err)
	assert.Equal(t, *created.Id, *found.Id)
	assert.Equal(t, *input.Name, *found.Name)
	assert.Equal(t, *input.Email, *found.Email)
	assert.Equal(t, input.UserFields["test"].(string), found.UserFields["test"].(string))

	searched, err := client.UserSearch(*input.Email)
	assert.NoError(t, err)
	assert.Len(t, searched, 1)
	assert.Equal(t, found, searched[0])
}
