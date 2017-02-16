package zendesk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationMembershipCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	assert.NoError(t, err)

	org1, err := randOrg(client)
	assert.NoError(t, err)
	org2, err := randOrg(client)
	assert.NoError(t, err)

	user, err := randUser(client)
	assert.NoError(t, err)

	// it should create an organization membership
	orgMembership1 := OrganizationMembership{
		UserID:         user.ID,
		OrganizationID: org1.ID,
	}

	created1, err := client.CreateOrganizationMembership(&orgMembership1)
	assert.NoError(t, err)
	assert.NotNil(t, created1.ID)
	assert.Equal(t, *user.ID, *created1.UserID)
	assert.Equal(t, *org1.ID, *created1.OrganizationID)

	orgMembership2 := OrganizationMembership{
		UserID:         user.ID,
		OrganizationID: org2.ID,
	}
	_, err = client.CreateOrganizationMembership(&orgMembership2)
	assert.NoError(t, err)

	// it should not throw error if existing membership is attempted to be created
	replayMembership, err := client.CreateOrganizationMembership(&orgMembership1)
	assert.NoError(t, err)
	assert.NotNil(t, replayMembership.ID)
	assert.Equal(t, *created1.UserID, *replayMembership.UserID)
	assert.Equal(t, *created1.OrganizationID, *replayMembership.OrganizationID)

	// it should return organization memberships for specific user
	found, err := client.ListOrganizationMembershipsByUserID(*user.ID)
	assert.NoError(t, err)
	assert.Len(t, found, 2)
}
