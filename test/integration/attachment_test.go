package integration

import (
	"os"
	"testing"

	"github.com/MEDIGO/go-zendesk/zendesk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttachmentCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	file, info := open(t, "ball.jpeg")

	// assert that it can upload a file
	upload, err := client.UploadFile("ball.jpeg", file)
	require.NoError(t, err)
	require.NotNil(t, upload.Token)
	require.NotNil(t, upload.Attachment)
	require.NotNil(t, upload.Attachment.ID)
	require.Equal(t, "ball.jpeg", *upload.Attachment.FileName)
	require.Equal(t, info.Size(), *upload.Attachment.Size)

	// assert that it can attach an upload to a ticket
	user, err := RandUser(client)
	assert.NoError(t, err)

	ticket, err := client.CreateTicket(&zendesk.Ticket{
		RequesterID: user.ID,
		Tags:        []string{"test"},
		Subject:     zendesk.String("My printer is on fire!"),
		Comment: &zendesk.TicketComment{
			Body:    zendesk.String("The smoke is very colorful."),
			Uploads: []string{*upload.Token},
		},
	})
	require.NoError(t, err)

	comments, err := client.ListTicketComments(*ticket.ID)
	require.NoError(t, err)
	require.Len(t, comments, 1)
	require.NotNil(t, comments[0].Attachments)

	attachments := comments[0].Attachments
	require.Len(t, attachments, 1)
	require.Equal(t, *upload.Attachment.ID, *attachments[0].ID)

}

func open(t *testing.T, name string) (*os.File, os.FileInfo) {
	file, err := os.Open("fixture/" + name)
	require.NoError(t, err)

	info, err := file.Stat()
	require.NoError(t, err)

	return file, info
}
