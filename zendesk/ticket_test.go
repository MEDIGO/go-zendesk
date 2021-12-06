package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTicketCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	user := randUser(t, client)
	defer client.DeleteUser(*user.ID)

	ticket := &Ticket{
		Subject:       String("My printer is on fire!"),
		Description:   String("The smoke is very colorful."),
		RequesterID:   user.ID,
		Tags:          []string{"test"},
		Collaborators: []interface{}{"email@example.com", &Collaborator{Name: String("NAME"), Email: String("email2@example.com")}},
	}

	created, err := client.CreateTicket(ticket)
	require.NoError(t, err)
	require.NotNil(t, created.ID)
	require.Len(t, created.Tags, 1)
	require.Len(t, created.CollaboratorIDs, 2)

	found, err := client.ShowTicket(*created.ID)
	require.NoError(t, err)
	require.Equal(t, created.Subject, found.Subject)
	require.Equal(t, created.RequesterID, found.RequesterID)
	require.Equal(t, created.Tags, found.Tags)

	colls, err := client.ListTicketCollaborators(*created.ID)
	require.NoError(t, err)
	require.Len(t, colls, 2)

	folls, err := client.ListTicketFollowers(*created.ID)
	require.NoError(t, err)
	require.Len(t, folls, 2)

	emailCCs, err := client.ListTicketEmailCCs(*created.ID)
	require.NoError(t, err)
	require.Len(t, emailCCs, 0)

	allTickets, err := client.ListTickets(&ListOptions{SortBy: "updated_at", SortOrder: "desc"}, IncludeUsers())
	require.NoError(t, err)
	require.True(t, len(allTickets.Tickets) > 0)
	require.Equal(t, *created.ID, *allTickets.Tickets[0].ID)

	safeUpdate := true
	input := Ticket{
		Status:                  String("solved"),
		AdditionalCollaborators: []interface{}{"email3@example.com"},
		UpdatedStamp:            created.UpdatedAt,
		SafeUpdate:              &safeUpdate,
	}

	updated, err := client.UpdateTicket(*created.ID, &input)
	require.NoError(t, err)
	require.Equal(t, input.Status, updated.Status)
	require.Len(t, updated.CollaboratorIDs, 3)

	requested, err := client.ListRequestedTickets(*user.ID)
	require.NoError(t, err)
	require.Len(t, requested, 1)
	require.Equal(t, created.ID, requested[0].ID)

	err = client.DeleteTicket(*created.ID)
	require.NoError(t, err)

	job, err := client.PermanentlyDeleteTicket(*created.ID)
	require.NoError(t, err)
	require.NotNil(t, job.ID)

	status, err := client.ShowJobStatus(*job.ID)
	require.NoError(t, err)
	require.NotNil(t, status.Status)

	externalID := "externalID" + randString(16)
	ticketReq := &Ticket{
		Subject:       String("My printer is on fire!"),
		Description:   String("The smoke is very colorful."),
		Tags:          []string{"test"},
		Collaborators: []interface{}{"email@example.com", &Collaborator{Name: String("NAME"), Email: String("email2@example.com")}},
		Requester:     &Requester{Email: String("email5@example.com"), Name: String("NAME2")},
		ExternalID:    &externalID,
	}

	createdReq, err := client.CreateTicket(ticketReq)
	require.NoError(t, err)
	require.NotNil(t, createdReq.ID)

	listRes, err := client.ListExternalIDTickets(externalID, nil)
	require.NoError(t, err)
	require.Equal(t, int64(1), *listRes.Count)
	require.Len(t, listRes.Tickets, 1)

	err = client.DeleteTicket(*createdReq.ID)
	require.NoError(t, err)

	jobReq, err := client.PermanentlyDeleteTicket(*createdReq.ID)
	require.NoError(t, err)
	require.NotNil(t, jobReq.ID)

}

func TestBatchUpdateManyTickets(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	user := randUser(t, client)
	defer client.DeleteUser(*user.ID)

	one := randTicket(t, client, user)
	defer client.DeleteTicket(*one.ID)

	two := randTicket(t, client, user)
	defer client.DeleteTicket(*two.ID)

	updates := []Ticket{
		{ID: one.ID, Status: String("solved")},
		{ID: two.ID, Status: String("solved")},
	}

	err = client.BatchUpdateManyTickets(updates)
	require.NoError(t, err)
}

func TestBulkUpdateManyTickets(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	user := randUser(t, client)
	defer client.DeleteUser(*user.ID)

	one := randTicket(t, client, user)
	defer client.DeleteTicket(*one.ID)

	two := randTicket(t, client, user)
	defer client.DeleteTicket(*two.ID)

	require.True(t, contains(one.Tags, "test"))
	require.True(t, contains(two.Tags, "test"))

	err = client.BulkUpdateManyTickets([]int64{*one.ID, *two.ID}, &Ticket{
		AdditionalTags: []string{"a_new_tag"},
		RemoveTags:     []string{"test"},
	})
	require.NoError(t, err)
}

func TestListTicketIncidents(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	user := randUser(t, client)
	defer client.DeleteUser(*user.ID)

	ticket := randTicket(t, client, user)
	defer client.DeleteTicket(*ticket.ID)

	incident1 := &Ticket{
		Description: String("Fire alarm trigger"),
		RequesterID: user.ID,
		Type:        String("incident"),
		ProblemID:   ticket.ID,
	}

	incident2 := &Ticket{
		Description: String("Building evacuation"),
		RequesterID: user.ID,
		Type:        String("incident"),
		ProblemID:   ticket.ID,
	}

	_, err = client.CreateTicket(incident1)
	require.NoError(t, err)
	_, err = client.CreateTicket(incident2)
	require.NoError(t, err)

	incidents, err := client.ListTicketIncidents(*ticket.ID)

	require.NoError(t, err)
	require.Len(t, incidents, 2)
}

func contains(l []string, s string) bool {
	for _, e := range l {
		if e == s {
			return true
		}
	}
	return false
}
