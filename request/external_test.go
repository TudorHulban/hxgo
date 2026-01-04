package request_test

import (
	"testing"

	"github.com/TudorHulban/hxgo/request"
	"github.com/stretchr/testify/require"
)

func TestRequest(t *testing.T) {
	formData, errParse := request.NewRequestData(
		&request.ParamsNewRequestData{
			BodyRaw:        []byte{},
			RequestHeaders: map[string][]string{},
		},
	)
	require.Error(t, errParse)
	require.Zero(t, formData)

	// _, _ = formData.ExtractMandatoryEpochID("some data")
}
