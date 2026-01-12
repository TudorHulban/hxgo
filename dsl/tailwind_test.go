package dsl

import (
	"fmt"
	"testing"

	"github.com/TudorHulban/hxgo/helpers"
)

func TestMethodsTailwind(t *testing.T) {
	c := TW()

	fmt.Println(
		helpers.MethodNamesOf(c),
	)
}

func TestTailwind(t *testing.T) {
	el := Div(
		Text("hi!"),

		TW().
			Absolute().
			FlexRow().
			AsNode(),
	)

	fmt.Println(
		string(Render(el)),
	)
}
