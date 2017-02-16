package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocaleCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	listed, err := client.ListLocales()
	require.NoError(t, err)
	require.NotEmpty(t, listed)

	found, err := client.ShowLocale(1)
	require.NoError(t, err)
	require.Equal(t, int64(1), *found.ID)

	found, err = client.ShowLocaleByCode("en-US")
	require.NoError(t, err)
	require.Equal(t, "en-US", *found.Locale)
}
