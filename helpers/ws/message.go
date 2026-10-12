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

	// RequestID is validated to [A-Za-z0-9_-]{1,64}, so it is safe to embed in
	// the HTML comment: <!-- _hx_req_id: {id} -->
	RequestID string
	Value     string

	IsPOST bool
}

func (m WSMessage) String() string {
	var sb strings.Builder

	sb.Grow(256)

	sb.WriteString("WSMessage{\n")
	fmt.Fprintf(&sb, "  Endpoint:  %q\n", m.Endpoint)
	fmt.Fprintf(&sb, "  IsPOST:    %t\n", m.IsPOST)
	fmt.Fprintf(&sb, "  CSRFToken: %q\n", m.CSRFToken)
	fmt.Fprintf(&sb, "  RequestID: %q\n", m.RequestID)
	fmt.Fprintf(&sb, "  Value:     %q\n", m.Value)

	if len(m.Values) > 0 {
		fmt.Fprintf(&sb, "  Values:    %s\n", m.Values.Encode())
	} else {
		sb.WriteString("  Values:    empty\n")
	}

	sb.WriteString("}")

	return sb.String()
}

func invalidInput(caller, raw string, issue error) error {
	return hxerrors.ErrInvalidInput{
		Issue:      issue,
		InputValue: raw,
		InputName:  "raw",
		Caller:     caller,
	}
}

// validRequestID allows only characters that cannot break out of an HTML comment.
func validRequestID(id string) bool {
	if len(id) > maxRequestIDLen {
		return false
	}

	for i := 0; i < len(id); i++ {
		c := id[i]

		switch {
		case c >= 'a' && c <= 'z',
			c >= 'A' && c <= 'Z',
			c >= '0' && c <= '9',
			c == '_', c == '-':

		default:
			return false
		}
	}

	return true
}

// parseFormMessage expects raw to start with "GET " or "POST "
// (guaranteed by ParseWSMessage). The body after the first newline is optional.
func parseFormMessage(raw string) (*WSMessage, error) {
	const caller = "parseFormMessage"

	head, body, _ := strings.Cut(raw, "\n")
	verb, endpoint, _ := strings.Cut(head, " ")

	endpoint = strings.TrimSpace(endpoint)
	if endpoint == "" {
		return nil,
			invalidInput(
				caller,
				raw,
				errors.New("malformed frame"),
			)
	}

	values, errParse := url.ParseQuery(body)
	if errParse != nil {
		return nil,
			invalidInput(
				caller,
				raw,
				fmt.Errorf("malformed body: %w", errParse),
			)
	}

	csrf := values.Get("_csrf")
	values.Del("_csrf")

	requestID := values.Get("_hx_req_id")
	values.Del("_hx_req_id")

	if !validRequestID(requestID) {
		return nil,
			invalidInput(
				caller,
				raw,
				errors.New("invalid request id"),
			)
	}

	return &WSMessage{
			Endpoint:  endpoint,
			CSRFToken: csrf,
			RequestID: requestID,
			Values:    values,
			IsPOST:    verb == "POST",
		},
		nil
}

func parsePipeMessage(raw string) (*WSMessage, error) {
	endpoint, value, couldCut := strings.Cut(raw, "|")
	if !couldCut || endpoint == "" {
		return nil,
			invalidInput(
				"parsePipeMessage",
				raw,
				errors.New("malformed frame"),
			)
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
