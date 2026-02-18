package components

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestARef(t *testing.T) {
	el := ARefRaw(
		&ParamsARef{
			Route:   "/home",
			Caption: "Home",
		},
	)

	require.Equal(t,
		`<a href="/home">Home</a>`,
		el,
	)
}
