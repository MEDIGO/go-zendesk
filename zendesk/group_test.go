package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	user := randUser(t, client)
	defer client.DeleteUser(*user.ID)

	group := &Group{
		Name: String("Test Group"),
	}

	beforeCreateAllGroups, err := client.ListGroups()
	require.NoError(t, err)

	created, err := client.CreateGroup(group)
	require.NoError(t, err)
	require.NotNil(t, created.ID)
	require.NotNil(t, created.URL)
	require.NotNil(t, created.CreatedAt)
	require.NotNil(t, created.UpdatedAt)

	afterCreateAllGroups, err := client.ListGroups()
	require.NoError(t, err)
	require.True(t, len(afterCreateAllGroups) > len(beforeCreateAllGroups))

	created.Name = String("Test")
	_, err = client.UpdateGroup(*created.ID, created)
	require.NoError(t, err)

	err = client.DeleteGroup(*created.ID)
	require.NoError(t, err)

	afterDeleteAllGroups, err := client.ListGroups()
	require.NoError(t, err)
	require.True(t, len(afterDeleteAllGroups) == len(beforeCreateAllGroups))
}
