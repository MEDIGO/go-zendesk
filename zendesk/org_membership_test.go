package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrganizationMembershipCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	org1 := randOrg(t, client)
	org2 := randOrg(t, client)

	user := randUser(t, client)

	// it should create an organization membership
	orgMembership1 := OrganizationMembership{
		UserID:         user.ID,
		OrganizationID: org1.ID,
	}

	created1, err := client.CreateOrganizationMembership(&orgMembership1)
	require.NoError(t, err)
	require.NotNil(t, created1.ID)
	require.Equal(t, *user.ID, *created1.UserID)
	require.Equal(t, *org1.ID, *created1.OrganizationID)

	orgMembership2 := OrganizationMembership{
		UserID:         user.ID,
		OrganizationID: org2.ID,
	}
	_, err = client.CreateOrganizationMembership(&orgMembership2)
	require.NoError(t, err)

	// it should not throw error if existing membership is attempted to be created
	replayMembership, err := client.CreateOrganizationMembership(&orgMembership1)
	require.NoError(t, err)
	require.NotNil(t, replayMembership.ID)
	require.Equal(t, *created1.UserID, *replayMembership.UserID)
	require.Equal(t, *created1.OrganizationID, *replayMembership.OrganizationID)

	// it should return organization memberships for specific user
	found, err := client.ListOrganizationMembershipsByUserID(*user.ID)
	require.NoError(t, err)
	require.Len(t, found, 2)
}
