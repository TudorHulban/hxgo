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

	Values url.Values
	IsPOST bool
}

func HandleWSMessage(raw string) (*WSMessage, error) {
	verbAndEndpoint, body, _ := strings.Cut(raw, "\n")

	verb, endpoint, couldCut := strings.Cut(verbAndEndpoint, " ")
	endpoint = strings.TrimSpace(endpoint)

	if !couldCut || len(endpoint) == 0 {
		return nil,
			hxerrors.ErrInvalidInput{
				Issue:      errors.New("malformed frame"),
				InputValue: raw,
				InputName:  "raw",
				Caller:     "HandleWSMessage",
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
				Caller:     "HandleWSMessage",
			}
	}

	csrf := values.Get("_csrf")
	values.Del("_csrf") // don't leak it into the handler's field map

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
			Caller:     "HandleWSMessage",
		}
}
