package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIdentityCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	// create user
	newUser := User{
		Name:  String(randString(16)),
		Email: String(randString(16) + "@example.com"),
	}

	user, err := client.CreateUser(&newUser)
	require.NoError(t, err)
	require.NotNil(t, user.ID)

	// get user
	_, err = client.ShowUser(*user.ID)
	require.NoError(t, err)

	// get user identity from list
	found, err := client.ListIdentities(*user.ID)
	require.NoError(t, err)
	require.Equal(t, len(found), 1)
	require.Equal(t, *found[0].Value, *user.Email)

	// update user identity
	newIdentity := UserIdentity{
		Value: String(randString(16) + "@example.com"),
	}

	_, err = client.UpdateIdentity(*user.ID, *found[0].ID, &newIdentity)
	require.NoError(t, err)

	// get user identity by id
	fetched, err := client.ShowIdentity(*user.ID, *found[0].ID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	require.Equal(t, fetched.Value, newIdentity.Value)

	// make sure the primary email has changed in user
	fetchedUser, err := client.ShowUser(*user.ID)
	require.NoError(t, err)
	require.NotNil(t, fetchedUser)
	require.Equal(t, *fetchedUser.Email, *newIdentity.Value)

	// create user identity (secondary email)
	secondaryIdentity := UserIdentity{
		Type:  String("email"),
		Value: String(randString(16) + "@example.com"),
	}

	created, err := client.CreateIdentity(*user.ID, &secondaryIdentity)
	require.NoError(t, err)

	found, err = client.ListIdentities(*user.ID)
	require.NoError(t, err)
	require.Equal(t, len(found), 2)

	// make second identity primary
	found, err = client.MakeIdentityPrimary(*user.ID, *created.ID)
	require.NoError(t, err)
	require.Equal(t, len(found), 2)

	// check if the list of identities shows the secondary identity is now primary
	for _, identity := range found {
		if *identity.ID == *created.ID {
			require.True(t, *identity.Primary)
		}
	}

	// delete user identity
	err = client.DeleteIdentity(*user.ID, *created.ID)
	require.NoError(t, err)

	found, err = client.ListIdentities(*user.ID)
	require.NoError(t, err)
	require.Equal(t, len(found), 1)
}
