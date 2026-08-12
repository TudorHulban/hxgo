package helpers

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/tudorhulban/hxerrors"
)

type WSMessage struct {
	Endpoint  string
	CSRFToken string

	Value  string
	Values url.Values
	IsPOST bool
}

func handleFormMessage(raw string) (*WSMessage, error) {
	verbAndEndpoint, body, couldCut := strings.Cut(raw, "\n")
	if !couldCut {
		return nil,
			hxerrors.ErrInvalidInput{
				Issue:      errors.New("malformed frame"),
				InputValue: raw,
				InputName:  "raw",
				Caller:     "handleFormMessage",
			}
	}

	verb, endpoint, _ := strings.Cut(verbAndEndpoint, " ")
	endpoint = strings.TrimSpace(endpoint)
	if len(endpoint) == 0 {
		return nil,
			hxerrors.ErrInvalidInput{
				Issue:      errors.New("malformed frame"),
				InputValue: raw,
				InputName:  "raw",
				Caller:     "handleFormMessage",
			}
	}

	values, errParse := url.ParseQuery(body)
	if errParse != nil {
		return nil,
			hxerrors.ErrInvalidInput{
				Issue: fmt.Errorf(
					"malformed body: %w",
					errParse,
				),
				InputValue: raw,
				InputName:  "raw",
				Caller:     "handleFormMessage",
			}
	}

	csrf := values.Get("_csrf")
	values.Del("_csrf")

	if verb == "POST" {
		return &WSMessage{
				Endpoint:  endpoint,
				CSRFToken: csrf,
				Values:    values,
				IsPOST:    true,
			},
			nil
	}

	if verb == "GET" {
		return &WSMessage{
				Endpoint:  endpoint,
				CSRFToken: csrf,
				Values:    values,
			},
			nil
	}

	return nil,
		hxerrors.ErrInvalidInput{
			Issue:      fmt.Errorf("unsupported verb %q", verb),
			InputValue: raw,
			InputName:  "raw",
			Caller:     "handleFormMessage",
		}
}

func handlePipeMessage(raw string) (*WSMessage, error) {
	endpoint, value, couldCut := strings.Cut(raw, "|")
	if !couldCut || len(endpoint) == 0 {
		return nil,
			hxerrors.ErrInvalidInput{
				Issue:      errors.New("malformed frame"),
				InputValue: raw,
				InputName:  "raw",
				Caller:     "handlePipeMessage",
			}
	}

	return &WSMessage{
			Endpoint: endpoint,
			Value:    value,
		},
		nil
}

func ParseWSMessage(raw string) (*WSMessage, error) {
	if strings.HasPrefix(raw, "GET ") || strings.HasPrefix(raw, "POST ") {
		return handleFormMessage(raw)
	}

	return handlePipeMessage(raw)
}
