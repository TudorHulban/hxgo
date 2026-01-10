package components

import (
	"fmt"
	"testing"

	"github.com/TudorHulban/hxgo/dsl"
	"github.com/stretchr/testify/require"
)

func TestNewForm3Containers(t *testing.T) {
	el := ButtonSubmitWCSS(
		&ParamsButtonSubmit{},

		dsl.NewCSSForClass("card:hover").
			WithBreakpoint("768px").
			Padding("10px 18px").
			BoxShadow("0 8px 24px rgba(0,0,0,0.2)").
			AsNode(),
	)

	form := NewFormThreeContainers(
		&ParamsNewFormThreeContainers{
			IDEnclosingDiv:  "footer-info",
			IDForm:          "form-footer",
			IDContainersDiv: "footer-form-info",

			ElementsInputLeft: []dsl.Node{
				dsl.Span(
					dsl.Text(
						"Contact and Open Hours",
					),
				),
				dsl.Span(
					dsl.Text(
						"+1234567890",
					),
				),
			},

			ElementsInputMiddle: []dsl.Node{
				dsl.Span(
					dsl.Text(
						"Address",
					),
				),
				dsl.Span(
					dsl.Text(
						"Lorem ipsum 3",
					),
				),
			},

			ElementsInputRight: []dsl.Node{
				dsl.Span(
					dsl.Text(
						"Book an appointment",
					),
				),

				el,
			},
		},
	)

	html, css := dsl.RenderHTMLandStyles(form)
	require.NotZero(t, html, "valid HTML")
	require.NotZero(t, css, "valid CSS")

	fmt.Println(
		string(html),
	)
	fmt.Println(css)
}
