package dsl

import (
	"fmt"
	"testing"

	"github.com/TudorHulban/hxgo/helpers"
	"github.com/tudorhulban/hxhelpers"
)

func TestMethodsTailwind(t *testing.T) {
	c := TW()

	if !hxhelpers.IsRunningInCI() {
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

	if !hxhelpers.IsRunningInCI() {
		fmt.Println(
			string(RenderFast(el)),
		)
	}
}
