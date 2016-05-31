package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MEDIGO/go-zendesk/zendesk"
)

func TestOrganizationCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	input := zendesk.Organization{
		Name: zendesk.String("test-" + RandString(7)),
		OrganizationFields: map[string]interface{}{
			"test": "this is a test",
		},
	}

	created, err := client.CreateOrganization(&input)
	assert.NoError(t, err)
	assert.NotNil(t, created.ID)
	assert.Equal(t, *input.Name, *created.Name)
	assert.Equal(t, input.OrganizationFields["test"].(string), "this is a test")

	found, err := client.ShowOrganization(*created.ID)
	assert.NoError(t, err)
	assert.Equal(t, *created.ID, *found.ID)
	assert.Equal(t, *input.Name, *found.Name)
	assert.Equal(t, input.OrganizationFields["test"].(string), found.OrganizationFields["test"].(string))
}

func TestOrganizationList(t *testing.T) {
	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	_, err = client.CreateOrganization(&zendesk.Organization{Name: zendesk.String("test-" + RandString(7))})
	_, err = client.CreateOrganization(&zendesk.Organization{Name: zendesk.String("test-" + RandString(7))})

	first, err := client.ListOrganizations(&zendesk.ListOptions{PerPage: 1})
	assert.NoError(t, err)
	assert.Len(t, first, 1)

	second, err := client.ListOrganizations(&zendesk.ListOptions{Page: 2, PerPage: 1})
	assert.NoError(t, err)
	assert.Len(t, first, 1)

	assert.NotEqual(t, *first[0].ID, *second[0].ID)
}
