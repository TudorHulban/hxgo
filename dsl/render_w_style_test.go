package dsl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStyledDiv(t *testing.T) {
	element := Div(
		GetStyledNode(
			Noop,
			[]Style{
				{
					Selector: ".card",
					Props: [][2]string{
						{"padding", "20px"},
						{"border-radius", "8px"},
						{"box-shadow", "0 4px 12px rgba(0,0,0,0.1)"},
					},

					Media: "min-width: 768px",
				},
				{
					Selector: ".card:hover",
					Props: [][2]string{
						{"box-shadow", "0 4px 12px rgba(0,0,0,0.1)"},
					},
				},
			}...,
		),

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
