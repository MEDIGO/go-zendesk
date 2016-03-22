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

	created, err := client.CreateOrganization(&input)
	assert.NoError(t, err)
	assert.NotNil(t, created.ID)
	assert.Equal(t, *input.Name, *created.Name)
	assert.Equal(t, input.OrganizationFields["test"].(string), "this is a test")

	found, err := client.GetOrganization(*created.ID)
	assert.NoError(t, err)
	assert.Equal(t, *created.ID, *found.ID)
	assert.Equal(t, *input.Name, *found.Name)
	assert.Equal(t, input.OrganizationFields["test"].(string), found.OrganizationFields["test"].(string))
}
