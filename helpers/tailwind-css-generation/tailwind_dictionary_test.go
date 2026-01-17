package tailwindcssgeneration

import (
	"fmt"
	"testing"

	"github.com/TudorHulban/hxgo/dsl"
	"github.com/stretchr/testify/require"
)

func TestDictionaryAligned(t *testing.T) {
	_ = dsl.TW().Absolute().
		AsNode()

	allMethods, errScan := ExtractAllMethods(dsl.TW())
	require.NoError(t, errScan)
	require.NotEmpty(t, allMethods)

	for _, methodInfo := range allMethods {
		_, exists := Dictionary[methodInfo.Name]
		require.True(
			t,
			exists,

			fmt.Sprintf(
				"searched dictionary (no methods = %d) - not found method %s",

				len(Dictionary),
				methodInfo.Name,
			),
		)
	}
}
