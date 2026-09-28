package request

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractMandatoryValueFromParentheses(t *testing.T) {
	data := RequestData{
		Content: map[string]string{
			"test_key": "example(value)another",
		},
	}

	expected := "value"

	output := data.ExtractMandatoryValueFromParentheses(
		"test_key",
	)
	require.Equal(t, expected, output)
}
