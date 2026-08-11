package main

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

type Server struct {
	app     *fiber.App
	counter atomic.Int64
	clients map[*websocket.Conn]bool
	mu      sync.RWMutex
}

func NewServer() *Server {
	s := &Server{
		clients: make(map[*websocket.Conn]bool),
	}

	app := fiber.New()

	app.Get(
		"/health",
		func(c fiber.Ctx) error {
			return c.SendString("OK")
		},
	)

	app.Use(
		"/ws",
		func(c fiber.Ctx) error {
			if c.Get("Upgrade") == "websocket" {
				return c.Next()
			}
			return fiber.ErrUpgradeRequired
		},
	)

	app.Get("/ws", websocket.New(s.handleWebSocket))

	app.Use("/", static.New("./public"))
	app.Use("/", static.New("../"))

	s.app = app

	return s
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

	handlers := map[string]func(*websocket.Conn, string){
		"/counter/increment": s.handleIncrement,
		"/counter/decrement": s.handleDecrement,
		"/counter/reset":     s.handleReset,
	}

	for {
		_, data, err := c.ReadMessage()
		if err != nil {
			break
		}

		msg := string(data)

		// Handle ping
		if msg == "ping" {
			c.WriteMessage(websocket.TextMessage, []byte("pong"))
			continue
		}

		// Parse route|value
		parts := strings.SplitN(msg, "|", 2)
		route := parts[0]

		handler, exists := handlers[route]
		if !exists {
			c.WriteMessage(websocket.TextMessage, []byte("unknown endpoint: "+route))
			continue
		}

		handler(c, route)
	}
}

func (s *Server) broadcast(html string) {
	s.mu.RLock()
	for c := range s.clients {
		c.WriteMessage(websocket.TextMessage, []byte(html))
	}
	s.mu.RUnlock()
}

func (s *Server) handleIncrement(c *websocket.Conn, route string) {
	val := s.counter.Add(1)

	html := fmt.Sprintf(`<div id="counter">%d</div>`, val)
	c.WriteMessage(websocket.TextMessage, []byte(html))

	s.broadcast(html)
}

func (s *Server) handleDecrement(c *websocket.Conn, route string) {
	val := s.counter.Add(-1)

	html := fmt.Sprintf(`<div id="counter">%d</div>`, val)
	c.WriteMessage(websocket.TextMessage, []byte(html))

	s.broadcast(html)
}

func (s *Server) handleReset(c *websocket.Conn, route string) {
	s.counter.Store(0)
	html := `<div id="counter">0</div>`
	c.WriteMessage(websocket.TextMessage, []byte(html))

	s.broadcast(html)
}

func (s *Server) Run(addr string) error {
	return s.app.Listen(addr)
}
