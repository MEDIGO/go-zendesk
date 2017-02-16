package zendesk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	assert.NoError(t, err)

	input := Organization{
		Name: String("test-" + randString(7)),
	}

	// it should create an organization
	created, err := client.CreateOrganization(&input)
	assert.NoError(t, err)
	assert.NotNil(t, created.ID)
	assert.Equal(t, *input.Name, *created.Name)

	// it should show an organization
	found, err := client.ShowOrganization(*created.ID)
	assert.NoError(t, err)
	assert.Equal(t, *created.ID, *found.ID)
	assert.Equal(t, *input.Name, *found.Name)

	name := "test-" + randString(7)

	// it should update an organization
	updated, err := client.UpdateOrganization(*found.ID, &Organization{
		Name: String(name),
	})

	assert.NoError(t, err)
	assert.Equal(t, name, *updated.Name)
}

func TestOrganizationList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	assert.NoError(t, err)

	_, err = client.CreateOrganization(&Organization{Name: String("test-" + randString(7))})
	_, err = client.CreateOrganization(&Organization{Name: String("test-" + randString(7))})

	first, err := client.ListOrganizations(&ListOptions{PerPage: 1})
	assert.NoError(t, err)
	assert.Len(t, first, 1)

	second, err := client.ListOrganizations(&ListOptions{Page: 2, PerPage: 1})
	assert.NoError(t, err)
	assert.Len(t, first, 1)

	assert.NotEqual(t, *first[0].ID, *second[0].ID)
}
