package zendesk

import (
	"crypto/rand"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func randString(l int) string {
	b := make([]byte, l)
	rand.Read(b)
	return hex.EncodeToString(b)[:l]
}

func randUser(t *testing.T, client Client) *User {
	input := &User{
		Name:  String("Testy Testacular"),
		Email: String(randString(16) + "@example.com"),
	}

	user, err := client.CreateUser(input)
	require.NoError(t, err)

	return user
}

func randOrg(t *testing.T, client Client) *Organization {
	input := &Organization{
		Name: String("Very Fake Clinic - " + randString(16)),
	}

	org, err := client.CreateOrganization(input)
	require.NoError(t, err)
	return org
}

func randTicket(t *testing.T, client Client, user *User) *Ticket {
	input := &Ticket{
		Subject:     String("My printer is on fire!"),
		Description: String("The smoke is very colorful."),
		RequesterID: user.ID,
		Type:        String("problem"),
		Tags:        []string{"test"},
	}

	ticket, err := client.CreateTicket(input)
	require.NoError(t, err)

	return ticket
}
