package integration

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/MEDIGO/go-zendesk/zendesk"
)

func TestOrganizationMembershipCRUD(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test in short mode.")
    }

    client, err := zendesk.NewEnvClient()
    assert.NoError(t, err)

    org, err := RandOrg(client)
    assert.NoError(t, err)

    user, err := RandUser(client)
    assert.NoError(t, err)
    
    // it should create an organization membership
    orgMembership := zendesk.OrganizationMembership{
        UserID:         user.ID,
        OrganizationID: org.ID,
    }

    created, err := client.CreateOrganizationMembership(&orgMembership)
    assert.NoError(t, err)
    assert.NotNil(t, created.ID)
    assert.Equal(t, *user.ID, *created.UserID)
    assert.Equal(t, *org.ID, *created.OrganizationID)
}