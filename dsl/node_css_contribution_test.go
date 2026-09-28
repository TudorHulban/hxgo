package dsl

import (
	"fmt"
	"testing"

	"github.com/TudorHulban/hxgo/helpers"
	"github.com/tudorhulban/hxhelpers"
)

func TestMethodsCSSContribution(t *testing.T) {
	c := CSSContribution{}

	if !hxhelpers.IsRunningInCI() {
		fmt.Println(
			len(
				helpers.MethodNamesOf(c),
			),
		)
	}
}
