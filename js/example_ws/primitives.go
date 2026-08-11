package main

import "github.com/gofiber/contrib/v3/websocket"

type Message struct {
	ID       string         `json:"id"`
	Verb     string         `json:"verb"`
	Endpoint string         `json:"endpoint"`
	Body     map[string]any `json:"body"`
	CSRF     string         `json:"csrf"`
}

type HandlerFunc func(*websocket.Conn, *Message)
