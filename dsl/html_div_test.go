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

func Test04DivStyles(t *testing.T) {
	cssClass := "css-class"

	el := Div(
		Class(cssClass),
		Text("hi!"),
	)

	el.Add(
		NewCSSForClass("card").
			WithBreakpoint("768px").
			Padding("10px 18px").
			AsNode(),

		NewCSSForClass("card").
			WithBreakpoint("1028px").
			Padding("15px 18px").
			AsNode(),
	)

	// elWithStyles x2 but only one style emitted.
	compound := Div(
		el,
		Text("-------"),
		el,
	)

	html, styles, css := RenderFull(compound)
	require.NotZero(t, html)
	require.NotZero(t, styles, "should have style")
	require.Zero(t, css)

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

func Test04aTailwind(t *testing.T) {
	el := Div(
		Class("bg-blue-500 p-4 rounded shadow"),
		Text("hi!"),
	)

	compound := Div(
		el,
	)

	html, styles, css := RenderFull(compound)
	require.NotZero(t, html)
	require.Zero(t, styles)
	require.Zero(t, css)

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

func Test05DivFull(t *testing.T) {
	cssClassComponent := "css-class"
	cssClassWidget := "css-widget"

	el := Div(
		Class(cssClassComponent),
		Class(cssClassWidget),
		Text("hi!"),
	)

	actualCSS := func() string {
		return `
			text-align: right;
			width: fit-content;
			background-color:rgb(134, 146, 138);
		`
	}

	el.Add(
		NewCSSForClass("card").
			WithBreakpoint("768px").
			Padding("10px 18px").
			AsNode(),

		NewCSSForClass("card").
			WithBreakpoint("1028px").
			Padding("15px 18px").
			AsNode(),

		NewCSSForClass(cssClassComponent).
			WithBreakpoint("1920px").
			AddCSS(actualCSS).
			AsNode(),
	)

	widget := Div(
		el,
		Text("-------"),
		el,
	)

	html, styles, css := RenderFull(widget)
	require.NotZero(t, html)
	require.NotZero(t, styles, "should have style")
	require.NotZero(t, css, "should have css")

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
	fmt.Println(
		"-----------------",
	)
	fmt.Println(
		css,
	)
}

func Test06DivCSS(t *testing.T) {
	el := Div(
		Text("hi!"),
	)

	actualCSS := func() string {
		return `
			text-align: right;
			width: fit-content;
			background-color:rgb(134, 146, 138);
		`
	}

	el.Add(
		NewGeneralCSS().
			AddCSS(actualCSS).
			AsNode(),
	)

	html, styles, css := RenderFull(el)
	require.NotZero(t, html)
	require.Zero(t, styles, "no styles")
	require.NotZero(t, css, "should have css")

	fmt.Println(
		string(html), // <div>hi!</div>
	)
	fmt.Println(
		"-----------------",
	)
	fmt.Println(
		css,
	)
}

func Test07DivCSS(t *testing.T) {
	el := Div(
		Text("hi!"),
	)

	actualCSS := func() string {
		return `
			text-align: right;
			width: fit-content;
			background-color:rgb(134, 146, 138);
		`
	}

	el.Add(
		NewGeneralCSS().
			WithBreakpoint("768px").
			AddCSS(actualCSS).
			AsNode(),
	)

	html, styles, css := RenderFull(el)
	require.NotZero(t, html)
	require.Zero(t, styles, "no styles")
	require.NotZero(t, css, "should have css")

	fmt.Println(
		string(html), // <div>hi!</div>
	)
	fmt.Println(
		"-----------------",
	)
	fmt.Println(
		css,
	)
}
