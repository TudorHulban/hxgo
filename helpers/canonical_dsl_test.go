package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanonical(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "01 empty input",
			input:       "",
			expected:    "",
		},
		{
			description: "02 whitespace outside quotes removed",
			input:       `  a  b  " c  d "  `,
			expected:    `ab" c d "`,
		},
		{
			description: "03 newlines removed",
			input:       "a\nb\nc",
			expected:    "abc",
		},
		{
			description: "04 trailing comma removed",
			input:       "f(a,b,)",
			expected:    "f(a,b)",
		},
		{
			description: "05 normalize parens",
			input:       "f ( a , b )",
			expected:    "f(a,b)",
		},
		{
			description: "06 collapse dot chains",
			input:       "a.b().c()",
			expected:    "a.b.c",
		},
		{
			description: "07 strip comments",
			input:       "a /* x */ b // y",
			expected:    "ab",
		},
		{
			description: "08 normalize comma spacing",
			input:       "a, b,  c",
			expected:    "a,b,c",
		},
		{
			description: "09 remove double spaces",
			input:       "a  b   c",
			expected:    "abc",
		},
		{
			description: "10 trim outer spaces",
			input:       "   abc   ",
			expected:    "abc",
		},
		{
			description: "11 canonical target dsl",
			input: `dsl.P(
        dsl.TW().TextSm().TextSmSm().TextSmMd().TextSmLg().TextSmXl().TextSmX2l(),
        dsl.Text("Hi!"),
    )`,
			expected: `P(TW.TextSm.TextSmSm.TextSmMd.TextSmLg.TextSmXl.TextSmX2l,Text("Hi!"))`,
		},
	}

	for _, test := range tests {
		t.Run(
			test.description,
			func(t *testing.T) {
				result := ApplyPredicates(test.input, CanonicalDSL...)
				require.Equal(t,
					test.expected,
					result,
				)
			},
		)
	}
}
