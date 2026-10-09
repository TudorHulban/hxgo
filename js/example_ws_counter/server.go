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

	result.serverWS.Register("/counter/increment", result.handleIncrement)
	result.serverWS.Register("/counter/decrement", result.handleDecrement)
	result.serverWS.Register("/counter/reset", result.handleReset)

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

func (s *Server) handleIncrement(c *ws.Client, message *ws.WSMessage) {
	html := fmt.Sprintf(
		`<div id="counter">%d</div>`,
		s.counter.Add(1),
	)

	if !c.Send(
		[]byte(
			ws.WrapResponse(message.RequestID, html),
		),
	) {
		log.Print(
			"write to client failed",
		)
	}

	s.serverWS.Broadcast([]byte(html), c)
}

func (s *Server) handleDecrement(c *ws.Client, message *ws.WSMessage) {
	html := fmt.Sprintf(
		`<div id="counter">%d</div>`,
		s.counter.Add(-1),
	)

	if !c.Send(
		[]byte(
			ws.WrapResponse(message.RequestID, html),
		),
	) {
		log.Print(
			"write to client failed",
		)
	}

	s.serverWS.Broadcast([]byte(html), c)
}

func (s *Server) handleReset(c *ws.Client, message *ws.WSMessage) {
	s.counter.Store(0)

	html := `<div id="counter">0</div>`

	if !c.Send(
		[]byte(
			ws.WrapResponse(message.RequestID, html),
		),
	) {
		log.Print(
			"write to client failed",
		)
	}

	s.serverWS.Broadcast([]byte(html), c)
}

func (s *Server) Run(addr string) error {
	return s.app.Listen(addr)
}
