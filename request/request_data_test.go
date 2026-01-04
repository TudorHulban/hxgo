package request

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractOptionalBoolValue(t *testing.T) {
	valueName := "blocked"

	reqData1 := RequestData{
		Content: map[string]string{
			valueName: "true",
		},
	}

	require.NotZero(t,
		reqData1.ExtractOptionalBoolValue(
			valueName,
		),
	)
}
