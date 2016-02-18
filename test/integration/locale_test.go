package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/medigo/go-zendesk/zendesk"
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

	found, err := client.LocaleGet(1)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), *found.ID)

	found, err = client.LocaleGetByCode("en-US")
	assert.NoError(t, err)
	assert.Equal(t, "en-US", *found.Locale)
}
