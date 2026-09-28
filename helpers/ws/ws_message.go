package ws

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/tudorhulban/hxerrors"
)

type WSMessage struct {
	Values url.Values

	Endpoint  string
	CSRFToken string

	RequestID string // response should include <!-- _hx_req_id: {id} -->
	Value     string

	IsPOST bool
}

func (m WSMessage) String() string {
	var sb strings.Builder

	// Pre-allocate memory to minimize allocations
	sb.Grow(256)

	fmt.Fprint(&sb, "WSMessage{\n")
	fmt.Fprintf(&sb,
		"  Endpoint:  %q\n",
		m.Endpoint,
	)
	fmt.Fprintf(&sb,
		"  IsPOST:    %t\n",
		m.IsPOST,
	)
	fmt.Fprintf(&sb,
		"  CSRFToken: %q\n",
		m.CSRFToken,
	)
	fmt.Fprintf(&sb,
		"  RequestID: %q (<!-- _hx_req_id: %s -->)\n",
		m.RequestID,
		m.RequestID,
	)
	fmt.Fprintf(&sb,
		"  Value:     %q\n",
		m.Value,
	)

	if len(m.Values) > 0 {
		fmt.Fprintf(&sb, "  Values:    %s\n", m.Values.Encode())
	} else {
		fmt.Fprint(&sb, "  Values:    empty\n")
	}

	sb.WriteString("}")

	return sb.String()
}

func parseFormMessage(raw string) (*WSMessage, error) {
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

	requestID := values.Get("_hx_req_id")
	values.Del("_hx_req_id")

	if verb == "POST" {
		return &WSMessage{
				Endpoint:  endpoint,
				CSRFToken: csrf,
				RequestID: requestID,

				Values: values,
				IsPOST: true,
			},
			nil
	}

	if verb == "GET" {
		return &WSMessage{
				Endpoint:  endpoint,
				CSRFToken: csrf,
				RequestID: requestID,

				Values: values,
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

func parsePipeMessage(raw string) (*WSMessage, error) {
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
		return parseFormMessage(raw)
	}

	return parsePipeMessage(raw)
}
