package request

import (
	"testing"

	goerrors "github.com/TudorHulban/go-errors"
	"github.com/stretchr/testify/require"
)

func TestGetNumberFromParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64

		expectInvalidInput bool
		exceptOtherError   bool
	}{
		{
			name:               "1_valid_number",
			input:              "example(42)another",
			expected:           42,
			expectInvalidInput: false,
		},
		{
			name:               "2_negative_number",
			input:              "value(-123)",
			expected:           -123,
			expectInvalidInput: false,
		},
		{
			name:               "3_zero_value",
			input:              "x(0)y",
			expected:           0,
			expectInvalidInput: false,
		},
		{
			name:               "4_only_opening_parenthesis",
			input:              "example(",
			expectInvalidInput: true,
		},
		{
			name:               "5_only_closing_parenthesis",
			input:              "example)",
			expectInvalidInput: true,
		},
		{
			name:               "6_no_parentheses",
			input:              "example",
			expectInvalidInput: true,
		},
		{
			name:               "7_multiple_parentheses_first_wins",
			input:              "example(10)another(20)",
			expected:           10,
			expectInvalidInput: false,
		},
		{
			name:               "8_empty_input",
			input:              "",
			expectInvalidInput: true,
		},
		{
			name:             "9_empty_value_in_parentheses",
			input:            "example()",
			exceptOtherError: true,
		},
		{
			name:             "10_non_numeric_value",
			input:            "example(abc)",
			exceptOtherError: true,
		},
	}

	for _, tc := range tests {
		t.Run(
			tc.name,
			func(t *testing.T) {
				result, err := getNumberFromParentheses(tc.input)

				if tc.expectInvalidInput {
					require.Error(t, err)

					var inputErr goerrors.ErrInvalidInput
					require.ErrorAs(t, err, &inputErr)
					require.Contains(t, err.Error(), "invalid input")

					return
				}

				if tc.exceptOtherError {
					require.Error(t, err)

					return
				}

				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			},
		)
	}
}

// BenchmarkGetNumberFromParentheses/1_simple_number-16         	31987947	       112.2 ns/op	       0 B/op	       0 allocs/op
// BenchmarkGetNumberFromParentheses/2_negative_number-16       	 8576008	       145.5 ns/op	       0 B/op	       0 allocs/op
// BenchmarkGetNumberFromParentheses/3_large_number-16          	 5935603	       206.4 ns/op	       0 B/op	       0 allocs/op
// BenchmarkGetNumberFromParentheses/4_multiple_parentheses-16  	 8108209	       134.5 ns/op	       0 B/op	       0 allocs/op
// BenchmarkGetNumberFromParentheses/5_empty_value-16           	 4510310	       256.6 ns/op	      48 B/op	       1 allocs/op
// BenchmarkGetNumberFromParentheses/6_no_parentheses-16        	 2255425	       509.5 ns/op	      96 B/op	       3 allocs/op

func BenchmarkGetNumberFromParentheses(b *testing.B) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "1_simple_number",
			input: "value(42)",
		},
		{
			name:  "2_negative_number",
			input: "x(-123)y",
		},
		{
			name:  "3_large_number",
			input: "n(9876543210)end",
		},
		{
			name:  "4_multiple_parentheses",
			input: "a(10)b(20)c",
		},
		{
			name:  "5_empty_value",
			input: "a()",
		},
		{
			name:  "6_no_parentheses",
			input: "abc",
		},
	}

	for _, tc := range tests {
		b.Run(
			tc.name,
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = getNumberFromParentheses(tc.input)
				}
			},
		)
	}
}
