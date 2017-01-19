package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MEDIGO/go-zendesk/zendesk"
)

func TestUserCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	input := zendesk.User{
		Name:  zendesk.String(RandString(16)),
		Email: zendesk.String(RandString(16) + "@example.com"),
	}

	created, err := client.CreateUser(&input)
	assert.NoError(t, err)
	assert.NotNil(t, created.ID)
	assert.Equal(t, *input.Name, *created.Name)
	assert.Equal(t, *input.Email, *created.Email)

	found, err := client.ShowUser(*created.ID)
	assert.NoError(t, err)
	assert.Equal(t, *created.ID, *found.ID)
	assert.Equal(t, *input.Name, *found.Name)
	assert.Equal(t, *input.Email, *found.Email)

	input = zendesk.User{
		Name: zendesk.String("Testy Testacular"),
	}

	updated, err := client.UpdateUser(*created.ID, &input)
	assert.NoError(t, err)
	assert.Equal(t, input.Name, updated.Name)

	searched, err := client.SearchUsers(*updated.Email)
	assert.NoError(t, err)
	assert.Len(t, searched, 1)
	assert.Equal(t, updated, &searched[0])

	other, err := client.CreateUser(&zendesk.User{
		Name:  zendesk.String(RandString(16)),
		Email: zendesk.String(RandString(16) + "@example.com"),
	})
	assert.NoError(t, err)

	input = zendesk.User{
		Name:  zendesk.String(RandString(16)),
		Email: updated.Email,
	}
	upserted, err := client.CreateOrUpdateUser(&input)
	assert.NoError(t, err)
	assert.NotNil(t, *upserted.ID)
	assert.Equal(t, *upserted.ID, *updated.ID)
	assert.NotEqual(t, *upserted.Name, *updated.Name)

	many, err := client.ShowManyUsers([]int64{*created.ID, *other.ID})
	assert.NoError(t, err)
	assert.Len(t, many, 2)
}

func TestListOrganizationUsers(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	org, err := client.CreateOrganization(&zendesk.Organization{
		Name: zendesk.String("test-" + RandString(7)),
	})
	assert.NoError(t, err)

	user, err := client.CreateUser(&zendesk.User{
		Name:           zendesk.String(RandString(16)),
		Email:          zendesk.String(RandString(16) + "@example.com"),
		OrganizationID: org.ID,
	})
	assert.NoError(t, err)

	found, err := client.ListOrganizationUsers(*org.ID, nil)
	assert.NoError(t, err)
	assert.Len(t, found, 1)
	assert.Equal(t, *user.ID, *found[0].ID)
}
