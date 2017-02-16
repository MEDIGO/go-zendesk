package zendesk

import (
	"crypto/rand"
	"encoding/hex"
)

func randString(l int) string {
	b := make([]byte, l)
	rand.Read(b)
	return hex.EncodeToString(b)[:l]
}

func randUser(client Client) (*User, error) {
	user := &User{
		Name:  String("Testy Testacular"),
		Email: String(randString(16) + "@example.com"),
	}

	return client.CreateUser(user)
}

func randOrg(client Client) (*Organization, error) {
	org := &Organization{
		Name: String("Very Fake Clinic - " + randString(16)),
	}

	return client.CreateOrganization(org)
}

func randTicket(client Client, user *User) (*Ticket, error) {
	ticket := &Ticket{
		Subject:     String("My printer is on fire!"),
		Description: String("The smoke is very colorful."),
		RequesterID: user.ID,
		Type:        String("problem"),
		Tags:        []string{"test"},
	}

	return client.CreateTicket(ticket)
}
