package dsl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStyledDiv(t *testing.T) {
	element := Div(
		Styled(
			nil,
			[]Style{
				{
					Selector: ".card",
					Props: map[string]string{
						"padding":       "20px",
						"border-radius": "8px",
						"box-shadow":    "0 4px 12px rgba(0,0,0,0.1)",
					},

					Media: "min-width: 768px",
				},
				{
					Selector: ".card:hover",
					Props: map[string]string{
						"box-shadow": "0 8px 24px rgba(0,0,0,0.2)",
					},
				},
			}...,
		),

		Div(AttrClass("card")),
	)

	html, css, errRender := RenderHTMLandCSS(element)
	require.NoError(t, errRender)
	require.NotZero(t, html)
	require.NotZero(t, css, "should have css")

	fmt.Println(
		string(
			html,
		),
	)
	fmt.Println(css)
}
