package ws

import (
	"sync"
	"time"

	"github.com/gofiber/contrib/v3/websocket"
)

// HandlerFunc handles one parsed message. It runs inline in the client's read
// loop, so a slow handler delays that client's subsequent messages (including
// pings). Reply with c.Send, never by writing to the raw connection.
type HandlerFunc func(c *Client, m *WSMessage)

type ServerWS struct {
	// handlers must be fully populated (via Handle) before the server accepts
	// connections; it is read without locking afterwards.
	handlers map[string]HandlerFunc

	clients map[*Client]struct{}
	mu      sync.RWMutex
}

// NewServer returns a server. A nil logger falls back to slog.Default().
func NewServer() *ServerWS {
	return &ServerWS{
		handlers: make(map[string]HandlerFunc),
		clients:  make(map[*Client]struct{}),
	}
}

// Handle registers fn for endpoint. Call before serving connections.
func (s *ServerWS) Register(endpoint string, fn HandlerFunc) {
	s.handlers[endpoint] = fn
}

// WrapResponse prefixes payload with the request-id marker.
// requestID is assumed to be validated by ParseWSMessage.
func WrapResponse(requestID, payload string) string {
	if requestID == "" {
		return payload
	}

	return reqIDPrefix + requestID + reqIDSuffix + payload
}

// WrapByteResponse is the []byte variant, with a single exact-size allocation.
func WrapByteResponse(requestID string, payload []byte) []byte {
	if requestID == "" {
		return payload
	}

	out := make([]byte, 0, len(reqIDPrefix)+len(requestID)+len(reqIDSuffix)+len(payload))
	out = append(out, reqIDPrefix...)
	out = append(out, requestID...)
	out = append(out, reqIDSuffix...)

	return append(out, payload...)
}

// Broadcast queues payload for every client except exclude (may be nil).
// It never blocks on the network. payload is shared between clients and must
// not be modified after the call.
func (s *ServerWS) Broadcast(payload []byte, exclude *Client) {
	s.mu.RLock()

	for c := range s.clients {
		if c != exclude {
			c.Send(payload)
		}
	}

	s.mu.RUnlock()
}

func (s *ServerWS) register(c *Client) {
	s.mu.Lock()
	s.clients[c] = struct{}{}
	s.mu.Unlock()
}

func (s *ServerWS) unregister(c *Client) {
	s.mu.Lock()
	delete(s.clients, c)
	s.mu.Unlock()
}

// HandleWebSocket is the fiber websocket handler. Enforce Origin checks in the
// upgrade middleware (websocket.Config{Origins: ...}).
func (s *ServerWS) HandleWebSocket(conn *websocket.Conn) {
	client := newClient(conn)

	s.register(client)

	go client.writeLoop()

	defer func() {
		s.unregister(client)
		client.Close()
	}()

	conn.SetReadLimit(_MaxMessageSize)

	extendDeadline := func() error {
		return conn.SetReadDeadline(time.Now().Add(_WaitPong))
	}

	_ = extendDeadline()

	conn.SetPongHandler(func(string) error { return extendDeadline() })

	for {
		_, raw, errReadMessage := conn.ReadMessage()
		if errReadMessage != nil {
			_ = websocket.IsUnexpectedCloseError( // s.log.Debug("websocket read ended", "err", err) if true
				errReadMessage,
				websocket.CloseNormalClosure,
				websocket.CloseGoingAway,
			)

			return
		}

		// Any inbound frame proves the peer is alive.
		_ = extendDeadline()

		s.dispatch(client, raw)
	}
}

func (s *ServerWS) dispatch(c *Client, raw []byte) {
	// Comparison against a constant does not allocate.
	if string(raw) == "ping" {
		c.Send(pongFrame)

		return
	}

	msg, errParse := ParseWSMessage(string(raw))
	if errParse != nil {
		// s.log.Debug("bad websocket frame", "err", errParse)
		c.Send(badRequestFrame)

		return
	}

	handler, exists := s.handlers[msg.Endpoint]
	if !exists {
		// s.log.Debug("unknown websocket endpoint", "endpoint", msg.Endpoint)
		c.Send(unknownEndpointFrame)

		return
	}

	s.safeCall(handler, c, msg)
}

// if r != nil:
// s.log.Error("websocket handler panic",
//
//	"endpoint", msg.Endpoint,
//	"panic", r,
//	"stack", string(debug.Stack()),
//
// )

// safeCall keeps a panicking handler from killing the connection goroutine.
func (*ServerWS) safeCall(handler HandlerFunc, c *Client, msg *WSMessage) {
	defer func() {
		if r := recover(); r != nil {
			c.Send(internalErrorFrame)
		}
	}()

	handler(c, msg)
}
