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
	defer client.DeleteOrganization(*org1.ID)

	org2 := randOrg(t, client)
	defer client.DeleteOrganization(*org2.ID)

	user := randUser(t, client)
	defer client.DeleteUser(*user.ID)

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
	created2, err := client.CreateOrganizationMembership(&orgMembership2)
	require.NoError(t, err)

	// it should return all organization memberships for specific user
	found, err := client.ListOrganizationMembershipsByUserID(*user.ID)
	require.NoError(t, err)
	require.Len(t, found, 2)
	found1 := isExistingMembership(*created1.UserID, *created1.OrganizationID, found)
	require.Equal(t, found1, true)
	found2 := isExistingMembership(*created2.UserID, *created2.OrganizationID, found)
	require.Equal(t, found2, true)

	// it should delete an organization membership
	err = client.DeleteOrganizationMembershipByID(*created1.ID)
	require.NoError(t, err)
	found, err = client.ListOrganizationMembershipsByUserID(*user.ID)
	require.NoError(t, err)
	require.Len(t, found, 1)
}

func isExistingMembership(userId, orgId int64, memberships []OrganizationMembership) bool {
	if memberships != nil {
		for _, membership := range memberships {
			if *membership.OrganizationID == orgId && *membership.UserID == userId {
				return true
			}
		}
	}

	return false
}
