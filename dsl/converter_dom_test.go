package dsl

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestTailwindTypographyMapping_PrintDSL(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		description string
		html        string
		expected    string
	}{
		{
			description: "1. transforms text-sm variants into DSL TW() chain",
			html: `
<p class="text-sm sm:text-sm md:text-sm lg:text-sm xl:text-sm 2xl:text-sm">
    Hi!
</p>`,
			expected: `dsl.P(
    dsl.TW().TextSm().TextSmSm().TextSmMd().TextSmLg().TextSmXl().TextSmX2l(),
    dsl.Text("Hi!"),
)
`,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.description,
			func(t *testing.T) {
				t.Parallel()

				doc, errParse := html.Parse(strings.NewReader(tc.html))
				require.NoError(t, errParse)

				root := ConvertHTML(doc)

				var buffer bytes.Buffer

				dom := ResultDOMConversion{
					Description: tc.description,
					Node:        root,
				}

				dom.PrintWithTransformers(
					&buffer,
					true,
					TailwindTransformer,
				)

				got := buffer.String()

				require.Contains(t,
					tc.expected,
					got,

					got,
				)
			},
		)
	}
}
