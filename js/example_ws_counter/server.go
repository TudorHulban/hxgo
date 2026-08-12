package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	"github.com/TudorHulban/hxgo/helpers"
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

type Server struct {
	app     *fiber.App
	counter atomic.Int64

	handlers map[string]func(*websocket.Conn, *helpers.WSMessage)

	clients map[*websocket.Conn]bool
	mu      sync.RWMutex
}

func NewServer() *Server {
	result := Server{
		clients: make(map[*websocket.Conn]bool),
	}
	result.handlers = map[string]func(*websocket.Conn, *helpers.WSMessage){
		"/counter/increment": result.handleIncrement,
		"/counter/decrement": result.handleDecrement,
		"/counter/reset":     result.handleReset,
	}

	app := fiber.New()

	app.Use(
		"/ws",
		func(c fiber.Ctx) error {
			if c.Get("Upgrade") == "websocket" {
				return c.Next()
			}

			return fiber.ErrUpgradeRequired
		},
	)

	app.Get("/ws", websocket.New(result.handleWebSocket))

	app.Use("/", static.New("./public"))
	app.Use("/", static.New("../"))

	result.app = app

	return &result
}

func (s *Server) handleWebSocket(c *websocket.Conn) {
	s.mu.Lock()
	s.clients[c] = true
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, c)
		s.mu.Unlock()
		c.Close()
	}()

	for {
		_, messageRaw, errRead := c.ReadMessage()
		if errRead != nil {
			break
		}

		messageString := string(messageRaw)

		// Handle ping
		if messageString == "ping" {
			c.WriteMessage(websocket.TextMessage, []byte("pong"))

			continue
		}

		wsMessage, errParse := helpers.ParseWSMessage(messageString)
		if errParse != nil {
			c.WriteMessage(websocket.TextMessage, []byte("bad request: "+errParse.Error()))

			continue
		}

		handler, exists := s.handlers[wsMessage.Endpoint]
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

func (s *Server) wrapResponse(requestID, html string) string {
	if requestID == "" {
		return html
	}

	return fmt.Sprintf(
		"<!-- _hx_req_id: %s -->\n%s",

		requestID,
		html,
	)
}

func (s *Server) broadcast(html string, exclude *websocket.Conn) {
	s.mu.RLock()
	for c := range s.clients {
		if c == exclude {
			continue
		}

		if errWrite := c.WriteMessage(
			websocket.TextMessage,
			[]byte(html),
		); errWrite != nil {
			log.Printf(
				"write to client failed: %v",
				errWrite,
			)
		}
	}
	s.mu.RUnlock()
}

func (s *Server) handleIncrement(c *websocket.Conn, message *helpers.WSMessage) {
	html := fmt.Sprintf(
		`<div id="counter">%d</div>`,
		s.counter.Add(1),
	)

	if errWrite := c.WriteMessage(
		websocket.TextMessage,
		[]byte(
			s.wrapResponse(message.RequestID, html),
		),
	); errWrite != nil {
		log.Printf(
			"write to client failed: %v",
			errWrite,
		)
	}

	s.broadcast(html, c)
}

func (s *Server) handleDecrement(c *websocket.Conn, message *helpers.WSMessage) {
	html := fmt.Sprintf(
		`<div id="counter">%d</div>`,
		s.counter.Add(-1),
	)

	if errWrite := c.WriteMessage(
		websocket.TextMessage,
		[]byte(
			s.wrapResponse(message.RequestID, html),
		),
	); errWrite != nil {
		log.Printf(
			"write to client failed: %v",
			errWrite,
		)
	}

	s.broadcast(html, c)
}

func (s *Server) handleReset(c *websocket.Conn, message *helpers.WSMessage) {
	s.counter.Store(0)

	html := `<div id="counter">0</div>`

	if errWrite := c.WriteMessage(
		websocket.TextMessage,
		[]byte(
			s.wrapResponse(message.RequestID, html),
		),
	); errWrite != nil {
		log.Printf(
			"write to client failed: %v",
			errWrite,
		)
	}

	s.broadcast(html, c)
}

func (s *Server) Run(addr string) error {
	return s.app.Listen(addr)
}
