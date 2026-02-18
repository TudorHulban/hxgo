package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsRunningInCI(t *testing.T) {
	require.False(t,
		IsRunningInCI(),
	)
}

func BenchmarkIsRunningInCI(b *testing.B) {
	for b.Loop() {
		IsRunningInCI()
	}
}
