package hxgo

import (
	"testing"

	"github.com/stretchr/testify/require"
	funcusage "github.com/tudorhulban/func-usage"
)

func TestFuncUsageAnalyzer(t *testing.T) {
	a, errCr := funcusage.NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	analysis, errAnalyze := a.Analyze(
		funcusage.ModeIncludeTestsForCoverage,
		true,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, analysis)

	printer := funcusage.NewPrinter().WithName()

	analysis.
		LevelFunction.
		WherePackageIs("dsl").
		IsFunction().
		AcceptingCaseInsensitiveLike("node").
		ReturningCaseInsensitiveLike("node").
		OrderByNameAsc().
		// MostUsed(10).
		PrintWith(printer)
}
