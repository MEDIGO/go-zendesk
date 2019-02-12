package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrganizationCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	input := Organization{
		Name: String("test-" + randString(7)),
	}

	// it should create an organization
	created, err := client.CreateOrganization(&input)
	require.NoError(t, err)
	require.NotNil(t, created.ID)
	require.Equal(t, *input.Name, *created.Name)

	// it should show an organization
	found, err := client.ShowOrganization(*created.ID)
	require.NoError(t, err)
	require.Equal(t, *created.ID, *found.ID)
	require.Equal(t, *input.Name, *found.Name)

	name := "test-" + randString(7)

	// it should update an organization
	updated, err := client.UpdateOrganization(*found.ID, &Organization{
		Name: String(name),
	})

	require.NoError(t, err)
	require.Equal(t, name, *updated.Name)

	err = client.DeleteOrganization(*created.ID)
	require.NoError(t, err)
}

func TestOrganizationCreateOrUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	input := Organization{
		Name: String("test-" + randString(7)),
	}

	// it should create an organization
	created, err := client.CreateOrUpdateOrganization(&input)
	require.NoError(t, err)
	require.NotNil(t, created.ID)
	require.Equal(t, *input.Name, *created.Name)

	// it should show an organization
	found, err := client.ShowOrganization(*created.ID)
	require.NoError(t, err)
	require.Equal(t, *created.ID, *found.ID)
	require.Equal(t, *input.Name, *found.Name)

	name := "test-" + randString(7)

	// it should update an organization
	updated, err := client.CreateOrUpdateOrganization(&Organization{
		ID:   found.ID,
		Name: String(name),
	})

	require.NoError(t, err)
	require.Equal(t, name, *updated.Name)

	err = client.DeleteOrganization(*created.ID)
	require.NoError(t, err)
}

func TestOrganizationList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	_, err = client.CreateOrganization(&Organization{Name: String("test-" + randString(7))})
	_, err = client.CreateOrganization(&Organization{Name: String("test-" + randString(7))})

	first, err := client.ListOrganizations(&ListOptions{PerPage: 1})
	require.NoError(t, err)
	require.Len(t, first, 1)

	second, err := client.ListOrganizations(&ListOptions{Page: 2, PerPage: 1})
	require.NoError(t, err)
	require.Len(t, first, 1)

	require.NotEqual(t, *first[0].ID, *second[0].ID)
}

func TestAutocompleteOrganizations(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	organizations, err := client.AutocompleteOrganizations(randString(7) + "very fake clinic")
	require.NoError(t, err)
	require.Len(t, organizations, 0)

	org1 := randOrg(t, client)
	defer client.DeleteOrganization(*org1.ID)

	organizations, err = client.AutocompleteOrganizations("very fake clinic")
	require.NoError(t, err)
	require.Len(t, organizations, 1)

	org2 := randOrg(t, client)
	defer client.DeleteOrganization(*org2.ID)

	organizations, err = client.AutocompleteOrganizations("very fake clinic")
	require.NoError(t, err)
	require.Len(t, organizations, 2)

}
