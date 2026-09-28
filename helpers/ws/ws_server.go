package ws

import (
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/contrib/v3/websocket"
)

// TODO: add logger.
type ServerWS struct {
	Handlers map[string]func(*websocket.Conn, *WSMessage)

	connections map[*websocket.Conn]bool
	mu          sync.RWMutex
}

func NewServer() *ServerWS {
	return &ServerWS{
		connections: make(map[*websocket.Conn]bool),
		Handlers:    make(map[string]func(*websocket.Conn, *WSMessage)),
	}
}

func (*ServerWS) WrapResponse(requestID, html string) string {
	if requestID == "" {
		return html
	}

	return fmt.Sprintf(
		"<!-- _hx_req_id: %s -->\n%s",

		requestID,
		html,
	)
}

func (s *ServerWS) Broadcast(html string, exclude *websocket.Conn) {
	s.mu.RLock()

	for conn := range s.connections {
		if conn == exclude {
			continue
		}

		if errWrite := conn.WriteMessage(
			websocket.TextMessage,
			[]byte(html),
		); errWrite != nil {
			log.Printf(
				"write to connection failed: %v",
				errWrite,
			)
		}
	}

	s.mu.RUnlock()
}

func (s *ServerWS) HandleWebSocket(c *websocket.Conn) {
	s.mu.Lock()
	s.connections[c] = true
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.connections, c)
		s.mu.Unlock()

		c.Close()
	}()

	for {
		_, messageRaw, errRead := c.ReadMessage()
		if errRead != nil {
			break
		}

		messageString := string(messageRaw)

		// fmt.Println(messageString)

		// Handle ping
		if messageString == "ping" {
			c.WriteMessage(
				websocket.TextMessage,
				[]byte("pong"),
			)

			continue
		}

		wsMessage, errParse := ParseWSMessage(messageString)
		if errParse != nil {
			c.WriteMessage(
				websocket.TextMessage,
				[]byte("bad request: "+errParse.Error()),
			)

			continue
		}

		// fmt.Println(wsMessage)

		handler, exists := s.Handlers[wsMessage.Endpoint]
		if !exists {
			c.WriteMessage(
				websocket.TextMessage,
				[]byte("unknown endpoint: "+wsMessage.Endpoint),
			)

			continue
		}

		handler(c, wsMessage)
	}
}
