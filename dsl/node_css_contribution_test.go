package dsl

import (
	"fmt"
	"testing"

	"github.com/TudorHulban/hxgo/helpers"
)

func TestMethodsCSSContribution(t *testing.T) {
	c := CSSContribution{}

	fmt.Println(
		helpers.MethodNamesOf(c),
	)
}
