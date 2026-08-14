package main

import (
	"fmt"
	"net/url"
	"strings"

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

	result.serverWS.Handlers["/login"] = result.wslogin

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

	result.app.Get(_RouteLogin, result.httpLogin)
	result.app.Get(_RouteAuthorised, result.authorized)
	result.app.Get(_RouteNotAuthorised, result.notauthorized)

	result.app.Use("/", static.New("./public"))
	result.app.Use("/", static.New("../"))

	return &result
}

func (s *Server) Run(addr string) error {
	return s.app.Listen(addr)
}

func extractCredentials(v url.Values) (string, string) {
	raw := v.Encode()

	var username, password string

	for _, pair := range strings.Split(raw, "&") {
		key, val, couldCut := strings.Cut(pair, "=")
		if !couldCut {
			continue
		}

		if key == "username" {
			username = val
		}

		if key == "password" {
			password = val
		}
	}

	return username, password
}

func (s *Server) wslogin(c *websocket.Conn, message *ws.WSMessage) {
	user, password := extractCredentials(message.Values)

	if user == "admin" && password == "password" {
		c.WriteMessage(
			websocket.TextMessage,
			[]byte(
				s.serverWS.WrapResponse(
					message.RequestID,
					fmt.Sprintf(
						`<div id="redirect" hx-redirect="%s"></div>`,
						_RouteAuthorised,
					),
				),
			),
		)
	}

	c.WriteMessage(
		websocket.TextMessage,
		[]byte(
			s.serverWS.WrapResponse(
				message.RequestID,
				fmt.Sprintf(
					`<div id="redirect" hx-redirect="%s"></div>`,
					_RouteNotAuthorised,
				),
			),
		),
	)
}

func (s *Server) httpLogin(c fiber.Ctx) error {
	return c.SendFile(
		"./public/index.html",
	)
}

func (s *Server) authorized(c fiber.Ctx) error {
	return c.SendFile(
		"./public/page_authorized.html",
	)
}

func (s *Server) notauthorized(c fiber.Ctx) error {
	return c.SendFile(
		"./public/page_not_authorized.html",
	)
}
