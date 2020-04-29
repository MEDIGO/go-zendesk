package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	input := User{
		Name:       String(randString(16)),
		Email:      String(randString(16) + "@example.com"),
		ExternalID: String(randString(16)),
	}

	created, err := client.CreateUser(&input)
	require.NoError(t, err)
	require.NotNil(t, created.ID)
	require.Equal(t, *input.Name, *created.Name)
	require.Equal(t, *input.Email, *created.Email)
	require.Equal(t, *input.ExternalID, *created.ExternalID)
	require.True(t, *created.Active)

	found, err := client.ShowUser(*created.ID)
	require.NoError(t, err)
	require.Equal(t, *created.ID, *found.ID)
	require.Equal(t, *input.Name, *found.Name)
	require.Equal(t, *input.Email, *found.Email)
	require.Equal(t, *input.ExternalID, *found.ExternalID)

	input = User{
		Name: String("Testy Testacular"),
	}

	updated, err := client.UpdateUser(*created.ID, &input)
	require.NoError(t, err)
	require.Equal(t, input.Name, updated.Name)

	searched, err := client.SearchUsers(*updated.Email)
	require.NoError(t, err)
	require.Len(t, searched, 1)
	require.Equal(t, updated, &searched[0])

	found, err = client.SearchUserByExternalID(*updated.ExternalID)
	require.NoError(t, err)
	require.Equal(t, updated, found)

	var nilUser *User
	found, err = client.SearchUserByExternalID("non-existent")
	require.NoError(t, err)
	require.Equal(t, nilUser, found)

	other, err := client.CreateUser(&User{
		Name:       String(randString(16)),
		Email:      String(randString(16) + "@example.com"),
		ExternalID: String(randString(16)),
	})
	require.NoError(t, err)

	input = User{
		Name:  String(randString(16)),
		Email: updated.Email,
	}
	upserted, err := client.CreateOrUpdateUser(&input)
	require.NoError(t, err)
	require.NotNil(t, *upserted.ID)
	require.Equal(t, *upserted.ID, *updated.ID)
	require.NotEqual(t, *upserted.Name, *updated.Name)

	many, err := client.ShowManyUsers([]int64{*created.ID, *other.ID})
	require.NoError(t, err)
	require.Len(t, many, 2)

	manyExternal, err := client.ShowManyUsersByExternalIDs([]string{*created.ExternalID, *other.ExternalID})
	require.NoError(t, err)
	require.Len(t, manyExternal, 2)

	tags, err := client.AddUserTags(*created.ID, []string{"a", "b"})
	require.NoError(t, err)
	require.Len(t, tags, 2)

	created, err = client.DeleteUser(*created.ID)
	require.NoError(t, err)
	require.False(t, *created.Active)

	statues, err := client.ShowComplianceDeletionStatuses(*created.ID)
	require.NoError(t, err)
	require.Zero(t, len(statues))

	_, err = client.PermanentlyDeleteUser(*created.ID)
	require.NoError(t, err)

	statues, err = client.ShowComplianceDeletionStatuses(*created.ID)
	require.NoError(t, err)
	require.NotZero(t, len(statues))

	other, err = client.DeleteUser(*other.ID)
	require.NoError(t, err)
	require.False(t, *other.Active)

	_, err = client.PermanentlyDeleteUser(*other.ID)
	require.NoError(t, err)

}

func TestListOrganizationUsers(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	org, err := client.CreateOrganization(&Organization{
		Name: String("test-" + randString(7)),
	})
	require.NoError(t, err)

	user, err := client.CreateUser(&User{
		Name:           String(randString(16)),
		Email:          String(randString(16) + "@example.com"),
		OrganizationID: org.ID,
	})
	require.NoError(t, err)

	found, err := client.ListOrganizationUsers(*org.ID, nil)
	require.NoError(t, err)
	require.Len(t, found, 1)
	require.Equal(t, *user.ID, *found[0].ID)
}

func TestListUsers(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	_, err = client.CreateUser(&User{
		Name:  String(randString(16)),
		Email: String(randString(16) + "@example.com"),
	})
	require.NoError(t, err)

	found, err := client.ListUsers(nil)
	require.NoError(t, err)
	require.NotEqual(t, 0, len(found))
}
