package dsl

import (
	"fmt"
	"testing"

	"github.com/TudorHulban/hxgo/helpers"
)

func TestMethodsCSSContribution(t *testing.T) {
	c := CSSContribution{}

	if !helpers.IsRunningInCI() {
		fmt.Println(
			len(
				helpers.MethodNamesOf(c),
			),
		)
	}
}
