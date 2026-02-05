package hxgo

import (
	"fmt"
	"testing"

	funcusage "github.com/TudorHulban/func-usage"
	"github.com/stretchr/testify/require"
)

func TestFuncUsageAnalyzer(t *testing.T) {
	a, errCr := funcusage.NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	usage, errAnalyze := a.Analyze(
		funcusage.ModeIncludeTestsForCoverage,
		true,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, usage)

	fmt.Println(
		usage.
			MostUsed(10).
			String(),
	)
}
