package request

import (
	"testing"

	goerrors "github.com/TudorHulban/go-errors"
	"github.com/stretchr/testify/require"
)

func TestGetValueFromParentheses(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    string
		expectError bool
	}{
		{
			description: "1_valid_input",
			input:       "example(value)another",
			expected:    "value",
			expectError: false,
		},
		{
			description: "2_only_opening_parenthesis",
			input:       "example(",
			expectError: true,
		},
		{
			description: "3_only_closing_parenthesis",
			input:       "example)",
			expectError: true,
		},
		{
			description: "4_no_parentheses",
			input:       "example",
			expectError: true,
		},
		{
			description: "5_multiple_parentheses",
			input:       "example(first)another(second)",
			expected:    "first",
			expectError: false,
		},
		{
			description: "6_empty_input",
			input:       "",
			expectError: true,
		},
		{
			description: "7_empty_value_in_parentheses",
			input:       "example()",
			expected:    "",
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(
			tc.description,
			func(t *testing.T) {
				result, err := getValueFromParentheses(tc.input)

				if tc.expectError {
					require.Error(t, err)

					var inputErr goerrors.ErrInvalidInput

					require.ErrorAs(t, err, &inputErr)
					require.Contains(t, err.Error(), "invalid input")

					return
				}

				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			},
		)
	}
}
