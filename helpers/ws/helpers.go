package ws

import "unsafe"

// WrapResponse prefixes payload with the request-id marker.
// requestID is assumed to be validated by ParseWSMessage.
func WrapResponse(requestID, payload string) string {
	if requestID == "" {
		return payload
	}

	return reqIDPrefix + requestID + reqIDSuffix + payload
}

// WrapResponses prefixes the concatenation of payloads with the request-id
// marker. With an empty requestID it just concatenates.
// requestID is assumed to be validated by ParseWSMessage.
func WrapResponses(requestID string, payloads ...string) string {
	if requestID == "" { // Zero-alloc fast path: nothing to wrap, nothing to join.
		switch len(payloads) {
		case 0:
			return ""
		case 1:
			return payloads[0]
		}
	}

	total := 0

	for ix := range payloads {
		total = total + len(payloads[ix])
	}

	if requestID != "" {
		total = total + len(reqIDPrefix) + len(requestID) + len(reqIDSuffix)
	}

	if total == 0 {
		return ""
	}

	buf := make([]byte, total) // single exact-size allocation, pointer-free (noscan)
	n := 0

	if requestID != "" {
		n = n + copy(buf[n:], reqIDPrefix)
		n = n + copy(buf[n:], requestID)
		n = n + copy(buf[n:], reqIDSuffix)
	}

	for ix := range payloads {
		n = n + copy(buf[n:], payloads[ix])
	}

	// buf never escapes or gets mutated after this point, so aliasing it as a
	// string is safe and avoids the copy that string(buf) would make.
	return unsafe.String(unsafe.SliceData(buf), total)
}

// WrapByteResponse is the []byte variant, with a single exact-size allocation.
func WrapByteResponse(requestID string, payload []byte) []byte {
	if requestID == "" {
		return payload
	}

	out := make(
		[]byte,
		0,
		len(reqIDPrefix)+len(requestID)+len(reqIDSuffix)+len(payload),
	)

	out = append(out, reqIDPrefix...)
	out = append(out, requestID...)
	out = append(out, reqIDSuffix...)

	return append(out, payload...)
}

func WrapByteResponses(requestID string, payloads ...[]byte) []byte {
	outputByteLength := 0

	for ix := range payloads {
		outputByteLength = outputByteLength + len(payloads[ix])
	}

	if requestID == "" {
		if len(payloads) == 1 {
			return payloads[0]
		}
	} else {
		outputByteLength = outputByteLength + len(reqIDPrefix) + len(requestID) + len(reqIDSuffix)
	}

	buf := make([]byte, outputByteLength)
	bufOffset := 0

	if requestID != "" {
		bufOffset = bufOffset + copy(buf[bufOffset:], reqIDPrefix)
		bufOffset = bufOffset + copy(buf[bufOffset:], requestID)
		bufOffset = bufOffset + copy(buf[bufOffset:], reqIDSuffix)
	}

	for ix := range payloads {
		bufOffset = bufOffset + copy(buf[bufOffset:], payloads[ix])
	}

	return buf
}
