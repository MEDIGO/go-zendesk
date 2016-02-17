package integration

import (
	"testing"

	"github.com/medigo/go-zendesk/zendesk"
	"github.com/stretchr/testify/assert"
)

func TestLocaleCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	listed, err := client.LocaleList()
	assert.NoError(t, err)
	assert.NotEmpty(t, listed)
}
