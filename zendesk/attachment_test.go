package zendesk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttachmentCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	assert.NoError(t, err)

	file, info := open(t, "ball.jpeg")

	// assert that it can upload a file
	upload1, err := client.UploadFile("ball.jpeg", nil, file)
	require.NoError(t, err)
	require.NotNil(t, upload1.Token)
	require.NotNil(t, upload1.Attachment)
	require.NotNil(t, upload1.Attachment.ID)
	require.Equal(t, "ball.jpeg", *upload1.Attachment.FileName)
	require.Equal(t, info.Size(), *upload1.Attachment.Size)

	file, info = open(t, "ball.jpeg") // reopen file

	// assert that it can reuse the upload token
	upload2, err := client.UploadFile("ball.jpeg", upload1.Token, file)
	require.NoError(t, err)

	// assert that it can attach the uploads to a ticket
	user, err := randUser(client)
	assert.NoError(t, err)

	ticket, err := client.CreateTicket(&Ticket{
		RequesterID: user.ID,
		Tags:        []string{"test"},
		Subject:     String("My printer is on fire!"),
		Comment: &TicketComment{
			Body:    String("The smoke is very colorful."),
			Uploads: []string{*upload1.Token},
		},
	})
	require.NoError(t, err)

	comments, err := client.ListTicketComments(*ticket.ID)
	require.NoError(t, err)
	require.Len(t, comments, 1)
	require.NotNil(t, comments[0].Attachments)

	attachments := comments[0].Attachments
	require.Len(t, attachments, 2)
	require.Equal(t, *upload1.Attachment.ID, *attachments[0].ID)
	require.Equal(t, *upload2.Attachment.ID, *attachments[1].ID)
}

func open(t *testing.T, name string) (*os.File, os.FileInfo) {
	file, err := os.Open("fixture/" + name)
	require.NoError(t, err)

	info, err := file.Stat()
	require.NoError(t, err)

	return file, info
}
