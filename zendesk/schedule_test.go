package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListSchedules(t *testing.T) {
	client, err := NewEnvClient()
	require.NoError(t, err)

	schedules, err := client.ListSchedules()
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(schedules), 1)
}
