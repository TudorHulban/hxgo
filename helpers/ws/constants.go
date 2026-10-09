package ws

import "time"

// Fixed replies: never echo client input back (it may be swapped into the DOM).
var (
	pongFrame            = []byte("pong")
	badRequestFrame      = []byte("bad request")
	unknownEndpointFrame = []byte("unknown endpoint")
	internalErrorFrame   = []byte("internal error")
)

const (
	reqIDPrefix = "<!-- _hx_req_id: "
	reqIDSuffix = " -->\n"
)

const (
	// _WaitWrite is the max time a single write may take before the client is dropped.
	_WaitWrite = 5 * time.Second

	// _WaitPong is how long we wait for any sign of life (message or pong) from the peer.
	_WaitPong = 60 * time.Second

	// _WaitPing must be less than pongWait.
	_WaitPing = (_WaitPong * 9) / 10

	// _MaxMessageSize caps inbound frames.
	_MaxMessageSize = 64 << 10

	// _SendQueuePerClient is the per-client outbound queue size.
	_SendQueuePerClient = 64
)

const maxRequestIDLen = 64
