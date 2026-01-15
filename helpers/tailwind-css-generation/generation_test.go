package tailwindcssgeneration

import (
	"os"
	"testing"

	"github.com/TudorHulban/hxgo/dsl"
	"github.com/stretchr/testify/require"
)

// This file itself contains the usage.

func TestOneLineScanDetectsAbsolute(t *testing.T) {
	_ = dsl.TW().Absolute().
		AsNode()

	allMethods, errScan := ExtractAllMethods(dsl.TW())
	require.NoError(t, errScan)
	require.NotEmpty(t, allMethods)

	usedMethods, errScan := ScanUsedMethods(
		ScanConfig{
			Files: []string{"./generation_test.go"},
		},
		allMethods,
	)
	require.NoError(t, errScan)
	require.Contains(t, usedMethods, "Absolute")
}

func TestMultiLineScanDetectsAbsolute(t *testing.T) {
	_ = dsl.TW().
		LeadingNormal().
		AsNode()

	allMethods, errScan := ExtractAllMethods(dsl.TW())
	require.NoError(t, errScan)
	require.NotEmpty(t, allMethods)

	usedMethods, errScan := ScanUsedMethods(
		ScanConfig{
			Files: []string{"./generation_test.go"},
		},
		allMethods,
	)
	require.NoError(t, errScan)
	require.Contains(t, usedMethods, "LeadingNormal")

	require.NoError(t,
		Writer(
			os.Stdout,
			Dictionary,
			usedMethods,
		),
	)
}
