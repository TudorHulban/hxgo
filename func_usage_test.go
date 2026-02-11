package hxgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	funcusage "github.com/tudorhulban/func-usage"
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

	usage.WherePackageIs("dsl")

	fmt.Println(
		usage.
			MostUsed(10).
			String(),
	)
}
