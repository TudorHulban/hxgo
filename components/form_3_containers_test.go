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
		[]dsl.Node{
			dsl.Styled(
				dsl.Noop,
				dsl.Style{
					Selector: ".card:hover",
					Props: [][2]string{
						{"box-shadow", "0 8px 24px rgba(0,0,0,0.2)"},
					},
				},
			),
		},
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

	html, css := dsl.RenderHTMLandCSS(form)
	require.NotZero(t, html, "valid HTML")
	require.NotZero(t, css, "valid CSS")

	fmt.Println(
		string(html),
	)
	fmt.Println(css)
}
