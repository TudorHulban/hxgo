package request_test

import (
	"testing"

	hxrequest "github.com/TudorHulban/hx-request"
	"github.com/stretchr/testify/require"
)

func TestRequest(t *testing.T) {
	formData, errParse := hxrequest.NewRequestData(
		&hxrequest.ParamsNewRequestData{
			BodyRaw:        []byte{},
			RequestHeaders: map[string][]string{},
		},
	)
	require.Error(t, errParse)
	require.Zero(t, formData)

	// _, _ = formData.ExtractMandatoryEpochID("some data")
}
