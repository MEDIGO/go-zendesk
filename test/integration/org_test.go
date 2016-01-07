package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/medigo/go-zendesk/zendesk"
)

func TestOrganizationCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	input := zendesk.Organization{
		Name: zendesk.String("test-" + randstr(7)),
		OrganizationFields: map[string]interface{}{
			"test": "this is a test",
		},
	}

	created, err := client.OrganizationCreate(&input)
	assert.NoError(t, err)
	assert.NotNil(t, created.Id)
	assert.Equal(t, *input.Name, *created.Name)
	assert.Equal(t, input.OrganizationFields["test"].(string), "this is a test")

	found, err := client.OrganizationGet(*created.Id)
	assert.NoError(t, err)
	assert.Equal(t, *created.Id, *found.Id)
	assert.Equal(t, *input.Name, *found.Name)
	assert.Equal(t, input.OrganizationFields["test"].(string), found.OrganizationFields["test"].(string))
}
