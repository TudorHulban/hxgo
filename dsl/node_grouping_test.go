package dsl

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProcBucket(t *testing.T) {
	f1 := func() string { return "a" }
	f2 := func() string { return "b" }

	b := newProcBucket()
	b.Add(f1)
	b.Add(f2)

	result := make([]string, 0, 2)

	for _, fn := range b.order {
		result = append(result, fn())
	}

	require.Equal(t,
		"ab",
		strings.Join(result, ""),
	)
}
