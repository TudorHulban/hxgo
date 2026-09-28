package hxgo

import (
	"testing"

	"github.com/stretchr/testify/require"
	funcusage "github.com/tudorhulban/func-usage"
)

func TestFuncUsageAnalyzer(t *testing.T) {
	t.Skipf(
		"test: '%s' is a manual test.",
		t.Name(),
	)

	a, errCr := funcusage.NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	analysis, errAnalyze := a.Analyze(
		funcusage.ModeIncludeTestsForCoverage,
		true,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, analysis)

	printer := funcusage.NewPrinter().
		WithName().
		WithPosition()

	analysis.
		LevelFunction.
		WherePackageIs("dsl").
		IsFunction().
		HasVariadic().
		AcceptingCaseInsensitiveLike("node").
		ReturningCaseInsensitiveLike("node").
		OrderByNameAsc().
		MostUsed(10).
		PrintWith(printer)
}
