package components

import (
	"fmt"
	"testing"
)

func TestARef(t *testing.T) {
	element := ARefRaw(
		&ParamsARef{
			Route:   "/home",
			Caption: "Home",
		},
	)

	fmt.Println(
		element,
	)
}
