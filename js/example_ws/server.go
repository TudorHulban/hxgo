package main

import (
	"encoding/json"
	"fmt"
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

	for {
		var msg struct {
			ID       string                 `json:"id"`
			Verb     string                 `json:"verb"`
			Endpoint string                 `json:"endpoint"`
			Body     map[string]interface{} `json:"body"`
			CSRF     string                 `json:"csrf"`
		}

		if err := c.ReadJSON(&msg); err != nil {
			break
		}

		if msg.Verb == "" && msg.Endpoint == "" {
			if msg.ID != "" {
				c.WriteJSON(map[string]string{"type": "pong"})
			}

			continue
		}

		broadcast := func(value int64) {
			response := map[string]any{
				"id":      msg.ID,
				"payload": fmt.Sprintf(`<div id="counter">%d</div>`, value),
			}
			c.WriteJSON(response)

			s.broadcast(value)
		}

		switch msg.Endpoint {
		case "/counter/increment":
			value := s.counter.Add(1)

			broadcast(value)

			fmt.Println(value)

		case "/counter/decrement":
			value := s.counter.Add(-1)

			broadcast(value)

			fmt.Println(value)

		case "/counter/reset":
			s.counter.Store(0)

			broadcast(0)

			fmt.Println(0)
		}

	}
}

func (s *Server) broadcast(val int64) {
	payload := fmt.Sprintf(`<div id="counter">%d</div>`, val)

	msg := map[string]any{
		"payload": payload,
	}
	data, _ := json.Marshal(msg)

	s.mu.RLock()
	for c := range s.clients {
		if err := c.WriteMessage(websocket.TextMessage, data); err != nil {
			c.Close()
		}
	}
	s.mu.RUnlock()
}

func (s *Server) Run(addr string) error {
	return s.app.Listen(addr)
}
