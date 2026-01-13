package tailwindcssgeneration

import (
	"testing"

	"github.com/TudorHulban/hxgo/dsl"
	"github.com/stretchr/testify/require"
)

// This file itself contains the usage.

func TestOneLineScanDetectsAbsolute(t *testing.T) {
	_ = dsl.TW().Absolute().
		AsNode()

	all, errScan := ExtractAllMethods(dsl.TW())
	require.NoError(t, errScan)
	require.NotEmpty(t, all)

	usedMethods, errScan := ScanUsedMethods(
		ScanConfig{
			Files: []string{"./generation_test.go"},
		},
		all,
	)
	require.NoError(t, errScan)
	require.Contains(t, usedMethods, "Absolute")
}

func TestMultiLineScanDetectsAbsolute(t *testing.T) {
	_ = dsl.TW().
		LeadingNormal().
		AsNode()

	all, errScan := ExtractAllMethods(dsl.TW())
	require.NoError(t, errScan)
	require.NotEmpty(t, all)

	usedMethods, errScan := ScanUsedMethods(
		ScanConfig{
			Files: []string{"./generation_test.go"},
		},
		all,
	)
	require.NoError(t, errScan)
	require.Contains(t, usedMethods, "LeadingNormal")
}
