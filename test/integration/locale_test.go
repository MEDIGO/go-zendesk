package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MEDIGO/go-zendesk/zendesk"
)

func TestLocaleCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := zendesk.NewEnvClient()
	assert.NoError(t, err)

	listed, err := client.ListLocales()
	assert.NoError(t, err)
	assert.NotEmpty(t, listed)

	found, err := client.GetLocale(1)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), *found.ID)

	found, err = client.GetLocaleByCode("en-US")
	assert.NoError(t, err)
	assert.Equal(t, "en-US", *found.Locale)
}
