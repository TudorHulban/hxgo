package components

import (
	"fmt"
	"os"
	"testing"

	"github.com/TudorHulban/hxgo/dsl"
)

func TestNewForm3Containers(t *testing.T) {
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
			},
		},
	)

	fmt.Println(
		form.Render(os.Stdout),
	)
}
