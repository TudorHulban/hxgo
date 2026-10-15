package ws

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrapResponse(t *testing.T) {
	response := WrapResponse("id3", "xxxxxxxxxxxxxxxx")

	expected := "<!-- _hx_req_id: id3 -->\nxxxxxxxxxxxxxxxx"

	require.Equal(t, expected, response)
}

func TestWrapResponses01(t *testing.T) {
	response := WrapResponses("id3", "xxxxxxx", "yyyyyyy")

	expected := "<!-- _hx_req_id: id3 -->\nxxxxxxx|yyyyyyy"

	require.Equal(t, expected, response)
}

func TestWrapResponses02(t *testing.T) {
	response := WrapResponses("id3", "x", "y", "z")

	expected := "<!-- _hx_req_id: id3 -->\nx|y|z"

	require.Equal(t, expected, response)
}

func TestWrapByteResponse(t *testing.T) {
	response := WrapByteResponse("id3", []byte("xxxxxxxxxxxxxxxx"))

	expected := "<!-- _hx_req_id: id3 -->\nxxxxxxxxxxxxxxxx"

	require.Equal(t, []byte(expected), response)
}

func TestWrapByteResponses(t *testing.T) {
	response := WrapByteResponses("id3", []byte("xxxxxxx"), []byte("yyyyyyy"))

	expected := "<!-- _hx_req_id: id3 -->\nxxxxxxx|yyyyyyy"

	require.Equal(t,
		[]byte(expected),
		response,
		string(response),
	)
}
