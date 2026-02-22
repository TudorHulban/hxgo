package dsl

import (
	"fmt"
	"testing"

	"github.com/TudorHulban/hxgo/helpers"
)

func TestMethodsTailwind(t *testing.T) {
	c := TW()

	if !helpers.IsRunningInCI() {
		fmt.Println(
			len(
				helpers.MethodNamesOf(c),
			),
		)
	}
}

func TestTailwind(t *testing.T) {
	el := Div(
		Text("hi!"),

		TW().
			Absolute().
			FlexRow().
			AsNode(),
	)

	if !helpers.IsRunningInCI() {
		fmt.Println(
			string(RenderFast(el)),
		)
	}
}
