package main

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/TudorHulban/hxgo/helpers/ws"
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

type Server struct {
	app      *fiber.App
	serverWS *ws.ServerWS

	counter atomic.Int64
}

func NewServer() *Server {
	result := Server{
		app:      fiber.New(),
		serverWS: ws.NewServer(),
	}

	result.serverWS.Handlers["/counter/increment"] = result.handleIncrement
	result.serverWS.Handlers["/counter/decrement"] = result.handleDecrement
	result.serverWS.Handlers["/counter/reset"] = result.handleReset

	result.app.Use(
		"/ws",
		func(c fiber.Ctx) error {
			if c.Get("Upgrade") == "websocket" {
				return c.Next()
			}

			return fiber.ErrUpgradeRequired
		},
	)

	result.app.Get("/ws", websocket.New(result.serverWS.HandleWebSocket))

	result.app.Use("/", static.New("./public"))
	result.app.Use("/", static.New("../"))

	return &result
}

func (s *Server) handleIncrement(c *websocket.Conn, message *ws.WSMessage) {
	html := fmt.Sprintf(
		`<div id="counter">%d</div>`,
		s.counter.Add(1),
	)

	if errWrite := c.WriteMessage(
		websocket.TextMessage,
		[]byte(
			s.serverWS.WrapResponse(message.RequestID, html),
		),
	); errWrite != nil {
		log.Printf(
			"write to client failed: %v",
			errWrite,
		)
	}

	s.serverWS.Broadcast(html, c)
}

func (s *Server) handleDecrement(c *websocket.Conn, message *ws.WSMessage) {
	html := fmt.Sprintf(
		`<div id="counter">%d</div>`,
		s.counter.Add(-1),
	)

	if errWrite := c.WriteMessage(
		websocket.TextMessage,
		[]byte(
			s.serverWS.WrapResponse(message.RequestID, html),
		),
	); errWrite != nil {
		log.Printf(
			"write to client failed: %v",
			errWrite,
		)
	}

	s.serverWS.Broadcast(html, c)
}

func (s *Server) handleReset(c *websocket.Conn, message *ws.WSMessage) {
	s.counter.Store(0)

	html := `<div id="counter">0</div>`

	if errWrite := c.WriteMessage(
		websocket.TextMessage,
		[]byte(
			s.serverWS.WrapResponse(message.RequestID, html),
		),
	); errWrite != nil {
		log.Printf(
			"write to client failed: %v",
			errWrite,
		)
	}

	s.serverWS.Broadcast(html, c)
}

func (s *Server) Run(addr string) error {
	return s.app.Listen(addr)
}
