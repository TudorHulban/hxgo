package dsl

import (
	"fmt"
	"testing"
)

func Test01Div(t *testing.T) {
	el := el("div", Text("hi!"))

	fmt.Println(
		string(Render(el)), // <div>hi!</div>
	)
}

func Test02Div(t *testing.T) {
	el := Div(
		Text("hi!"),
	)

	fmt.Println(
		string(Render(el)), // <div>hi!</div>
	)
}

func Test03Div(t *testing.T) {
	el := Div(
		Text("hi"),
		Div(
			Span(
				Text("!"),
			),
		),
	)

	fmt.Println(
		string(Render(el)), // <div>hi<div><span>!</span></div></div>
	)
}
