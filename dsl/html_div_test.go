package dsl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test01Div(t *testing.T) {
	el := el("div", Text("hi!"))

	require.Equal(t,
		"<div>hi!</div>",
		string(Render(el)),

		string(Render(el)),
	)
}

func Test02Div(t *testing.T) {
	el := Div(
		Class("css-class"),
		Text("hi!"),
	)

	require.Equal(t,
		`<div class="css-class">hi!</div>`,
		string(Render(el)),

		string(Render(el)),
	)
}

func Test03Div(t *testing.T) {
	el := Div(
		Class("css-class"),
		Text("hi"),

		Div(
			Span(
				Text("!"),
			),
		),
	)

	require.Equal(t,
		`<div class="css-class">hi<div><span>!</span></div></div>`,
		string(Render(el)),

		string(Render(el)),
	)
}

func Test04Div(t *testing.T) {
	cssClass := "css-class"

	el := Div(
		Class(cssClass),
		Text("hi!"),
	)

	elWithStyles := GetStyledNode(
		el,
		Style{
			Selector: "." + cssClass,
			Props: [][2]string{
				{"padding", "10px 18px"},
			},
			Media: "768px",
		},
		Style{
			Selector: "." + cssClass,
			Props: [][2]string{
				{"padding", "15px 18px"},
			},
			Media: "1028px",
		},
	)

	// elWithStyles x2 but only one style emitted.
	compound := Div(
		elWithStyles,
		Text("-------"),
		elWithStyles,
	)

	html, styles := RenderHTMLandStyles(compound)
	require.NotZero(t, html)
	require.NotZero(t, styles, "should have style")

	// <div><div class="css-class">hi!</div>-------<div class="css-class">hi!</div></div>
	//   .css-class {
	//     padding: 10px 18px;
	//   }

	fmt.Println(
		string(html),
	)
	fmt.Println(
		styles,
	)
}

func Test05Div(t *testing.T) {
	cssClassComponent := "css-class"
	cssClassWidget := "css-widget"

	el := Div(
		Class(cssClassComponent),
		Class(cssClassWidget),
		Text("hi!"),
	)

	elWithStyles := GetStyledNode(
		el,
		Style{
			Selector: "." + cssClassComponent,
			Props: [][2]string{
				{"padding", "10px 18px"},
			},
			Media: "768px",
		},
		Style{
			Selector: "." + cssClassComponent,
			Props: [][2]string{
				{"padding", "15px 18px"},
			},
			Media: "1028px",
		},
	)

	actualCSS := func() string {
		return `
		.css-widget {
			text-align: right;
			width: fit-content;
			background-color:rgb(134, 146, 138);

			#resource-selection {
				padding-top: 2.1em;
			}

			.hours-grid {
				width: 100%;
			}
		}
		`
	}

	cssRule := CSS{
		CSSKey: CSSKey{
			Selector:         cssClassWidget,
			InflexionPointPX: 768,
		},

		ActualCSS: actualCSS,
	}

	elWithCSS := GetCSSNode(cssRule)

	// elWithStyles x2 but only one style emitted.
	widget := Div(
		elWithStyles,
		Text("-------"),
		elWithStyles,
		elWithCSS,
	)

	html, styles := RenderHTMLandStyles(widget)
	require.NotZero(t, html)
	require.NotZero(t, styles, "should have style")

	// <div><div class="css-class">hi!</div>-------<div class="css-class">hi!</div></div>
	//   .css-class {
	//     padding: 10px 18px;
	//   }

	fmt.Println(
		string(html),
	)
	fmt.Println(
		styles,
	)
}
