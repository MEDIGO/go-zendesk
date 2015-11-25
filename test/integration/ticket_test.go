package integration

import (
	"testing"

	"github.com/medigo/go-zendesk/zendesk"
	"github.com/stretchr/testify/assert"
)

func TestTicketCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	_, err := zendesk.NewEnvClient()
	assert.NoError(t, err)
}
