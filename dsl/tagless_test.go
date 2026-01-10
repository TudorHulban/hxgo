package dsl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImg(t *testing.T) {
	el := Img()

	require.Equal(t,
		`<img>`,
		string(Render(el)),

		string(Render(el)),
	)
}
