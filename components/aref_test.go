package components

import (
	"fmt"
	"testing"
)

func TestARef(t *testing.T) {
	el := ARefRaw(
		&ParamsARef{
			Route:   "/home",
			Caption: "Home",
		},
	)

	fmt.Println(
		el,
	)
}
