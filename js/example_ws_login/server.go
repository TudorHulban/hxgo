package main

import (
	"github.com/TudorHulban/hxgo/helpers/ws"
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

type Server struct {
	app      *fiber.App
	serverWS *ws.ServerWS
}

func NewServer() *Server {
	result := Server{
		app:      fiber.New(),
		serverWS: ws.NewServer(),
	}

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

func (s *Server) Run(addr string) error {
	return s.app.Listen(addr)
}
