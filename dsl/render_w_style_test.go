package dsl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStyledDiv(t *testing.T) {
	element := Div(
		CSSContribution{
			CSSContributionKey: CSSContributionKey{
				Selector:       ".card",
				InflexionPoint: "768px",
			},
			DeclarativeStyle: [][2]string{
				{"padding", "20px"},
				{"border-radius", "8px"},
				{"box-shadow", "0 4px 12px rgba(0,0,0,0.1)"},
			},
		}.
			AsNode(),

		CSSContribution{
			CSSContributionKey: CSSContributionKey{
				Selector: ".card",
			},
			DeclarativeStyle: [][2]string{
				{"box-shadow", "0 4px 12px rgba(0,0,0,0.1)"},
			},
		}.
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
