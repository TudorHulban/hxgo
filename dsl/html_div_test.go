package dsl

import (
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
