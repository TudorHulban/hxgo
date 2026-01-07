package dsl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStyledDiv(t *testing.T) {
	element := Div(
		NewCSSForClass("card").
			WithBreakpoint("768px").
			Padding("20px").
			Radius("8px").
			ShadowBox("0 4px 12px rgba(0,0,0,0.1)").
			AsNode(),

		NewCSSForClass("card").
			WithBreakpoint("1024px").
			ShadowBox("0 4px 12px rgba(0,0,0,0.1)").
			AsNode(),

		Div(AttrClass("card")),
	)

	html, css := RenderHTMLandStyles(element)
	require.NotZero(t, html)
	require.NotZero(t, css, "should have css")

	fmt.Println(
		string(
			html,
		),
	)
	fmt.Println(css)
}
