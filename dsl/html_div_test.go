package dsl

import (
	"fmt"
	"testing"

	"github.com/TudorHulban/hxgo/helpers"
	"github.com/stretchr/testify/require"
)

func Test01Div(t *testing.T) {
	el := el("div", Text("hi!"))

	require.Equal(t,
		"<div>hi!</div>",
		string(RenderFast(el)),

		string(RenderFast(el)),
	)
}

func Test02Div(t *testing.T) {
	el := Div(
		Class("css-class"),
		Text("hi!"),
	)

	require.Equal(t,
		`div[class="css-class"](text("hi!"))`,
		el.Canonical(),
	)

	require.Equal(t,
		`<div class="css-class">hi!</div>`,
		string(RenderFast(el)),

		string(RenderFast(el)),
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
		el.Canonical(),
		`div[class="css-class"](text("hi"), div(span(text("!"))))`,
	)

	require.Equal(t,
		`<div class="css-class">hi<div><span>!</span></div></div>`,
		string(RenderFast(el)),

		string(RenderFast(el)),
	)
}

func Test04DivStyle(t *testing.T) {
	el := Div(
		Class("css-class"),
		Text(
			fmt.Sprintf(
				"hi %s!",
				t.Name(),
			),
		),
	)

	type TData struct {
		tag      string
		attrs    []Node
		children []Node
	}

	require.Len(t,
		(*TData)(el.data).attrs,
		1,
	)
	require.Len(t,
		(*TData)(el.data).children,
		1,
	)

	el.Add(
		NewCSSForClass("card").
			WithBreakpoint("768px").
			Padding("10px 10px").
			AsNode(),

		NewCSSForClass("card").
			WithBreakpoint("1028px").
			Padding("18px 18px").
			AsNode(),
	)

	require.Len(t,
		(*TData)(el.data).children,
		1,
	)
	require.Len(t,
		el.children,
		2,
	)

	t.Run(
		"1. Div direct",
		func(t *testing.T) {
			html, styles, css := RenderFull(el)
			require.Zero(t, css)
			require.NotZero(t, html)
			require.NotZero(t, styles, "should have style")

			if !helpers.IsRunningInCI() {
				fmt.Println(
					string(html),
				)
				fmt.Println(
					styles,
				)
			}
		},
	)

	t.Run(
		"2. Div soup",
		func(t *testing.T) {
			compound := Div(el)

			html, styles, css := RenderFull(compound)
			require.Zero(t, css)
			require.NotZero(t, html)
			require.NotZero(t, styles, "should have style")
		},
	)

	t.Run(
		"3. Unique CSS emitted",
		func(t *testing.T) {
			compound := Div(
				el,
				el,
			)

			html, styles, css := RenderFull(compound)
			require.Zero(t, css)
			require.NotZero(t, html)
			require.NotZero(t, styles, "should have style")

			if !helpers.IsRunningInCI() {
				fmt.Println(
					string(html),
				)
				fmt.Println(
					styles,
				)
			}
		},
	)
}

func Test06Tailwind(t *testing.T) {
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

	if !helpers.IsRunningInCI() {
		fmt.Println(
			string(html),
		)
		fmt.Println(
			styles,
		)
	}
}

func Test07DivFull(t *testing.T) {
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

	if !helpers.IsRunningInCI() {
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
}

func Test08DivCSS(t *testing.T) {
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

	if !helpers.IsRunningInCI() {
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
}

func Test09DivCSS(t *testing.T) {
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

	if !helpers.IsRunningInCI() {
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
}
