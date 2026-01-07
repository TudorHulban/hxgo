package dsl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAllocGuard_EmptyInput(t *testing.T) {
	const maxAllocs = 0

	allocs := testing.AllocsPerRun(
		1000,
		func() {
			_, _ = RenderHTMLandStyles()
		},
	)

	require.LessOrEqual(t,
		allocs,
		float64(maxAllocs),

		"alloc regression: got %.2f allocs/op, want <= %d", allocs, maxAllocs,
	)
}

func TestAllocGuard_HTMLOnly(t *testing.T) {
	// Adjust this ceiling based on your current benchmark.
	// After HTML fast-path optimization, this should drop to 2–3.
	const maxAllocs = 5

	node := Div(Class("card"), Text("hi!"))

	allocs := testing.AllocsPerRun(
		1000,
		func() {
			_, _ = RenderHTMLandStyles(node)
		},
	)

	require.LessOrEqual(t,
		allocs,
		float64(maxAllocs),

		"alloc regression: got %.0f allocs/op, want <= %d", allocs, maxAllocs,
	)
}

func TestAllocGuard_CSSOnly(t *testing.T) {
	const maxAllocs = 16

	node := NewCSSForClass("card").
		WithBreakpoint("768px").
		Padding("20px").
		Radius("8px").
		ShadowBox("0 4px 12px rgba(0,0,0,0.1)").
		AsNode()

	allocs := testing.AllocsPerRun(
		200,
		func() {
			_, _ = RenderHTMLandStyles(node)
		},
	)

	require.LessOrEqual(t,
		allocs,
		float64(maxAllocs),

		"alloc regression: got %.0f allocs/op, want <= %d", allocs, maxAllocs,
	)
}

func TestAllocGuard_HTMLAndCSS(t *testing.T) {
	const maxAllocations = 9

	node := Div(
		Class("card"),
		NewCSSForClass("card").
			WithBreakpoint("768px").
			Padding("20px").
			Radius("8px").
			ShadowBox("0 4px 12px rgba(0,0,0,0.1)").
			AsNode(),
	)

	allocs := testing.AllocsPerRun(
		200,
		func() {
			_, _ = RenderHTMLandStyles(node)
		},
	)

	require.LessOrEqual(t,
		allocs,
		float64(maxAllocations),

		"alloc regression: got %.0f allocs/op, want <= %d", allocs, maxAllocations,
	)
}
