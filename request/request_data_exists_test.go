package request

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEntryExists(t *testing.T) {
	data := RequestData{
		Content: map[string]string{
			"test_key": "(123)",
		},
	}

	require.False(t,
		data.EntryExists("test"),
		"T1. Entry does not exist.",
	)

	require.True(t,
		data.EntryExists("test_key"),
		"T2. Entry exists.",
	)
}
