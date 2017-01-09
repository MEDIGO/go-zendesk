package integration

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/MEDIGO/go-zendesk/zendesk"
)

func RandString(l int) string {
	b := make([]byte, l)
	rand.Read(b)
	return hex.EncodeToString(b)[:l]
}

func RandUser(client zendesk.Client) (*zendesk.User, error) {
	user := &zendesk.User{
		Name:  zendesk.String("Testy Testacular"),
		Email: zendesk.String(RandString(16) + "@example.com"),
	}

	return client.CreateUser(user)
}

func RandOrg(client zendesk.Client) (*zendesk.Organization, error) {
	org := &zendesk.Organization{
		Name:  zendesk.String("Very Fake Clinic - " + RandString(16)),
	}

	return client.CreateOrganization(org)
}

func RandTicket(client zendesk.Client, user *zendesk.User) (*zendesk.Ticket, error) {
	ticket := &zendesk.Ticket{
		Subject:     zendesk.String("My printer is on fire!"),
		Description: zendesk.String("The smoke is very colorful."),
		RequesterID: user.ID,
		Type:	     zendesk.String("problem"),
		Tags:        []string{"test"},
	}

	return client.CreateTicket(ticket)
}
