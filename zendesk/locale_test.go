package zendesk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocaleCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	assert.NoError(t, err)

	listed, err := client.ListLocales()
	assert.NoError(t, err)
	assert.NotEmpty(t, listed)

	found, err := client.ShowLocale(1)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), *found.ID)

	found, err = client.ShowLocaleByCode("en-US")
	assert.NoError(t, err)
	assert.Equal(t, "en-US", *found.Locale)
}
