package ws

import (
	"sync"
	"time"

	"github.com/gofiber/contrib/v3/websocket"
)

// Client owns a websocket connection. All writes go through send and are
// performed by a single goroutine (writeLoop), which satisfies the
// one-concurrent-writer rule of the underlying connection.
type Client struct {
	conn *websocket.Conn

	chSend chan []byte
	chDone chan struct{}

	closeOnce sync.Once
}

func newClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,

		chSend: make(chan []byte, _SendQueuePerClient),
		chDone: make(chan struct{}),
	}
}

// Conn exposes the raw connection for read-only uses such as Locals().
// Do not write to it directly.
func (c *Client) Conn() *websocket.Conn {
	return c.conn
}

// Close is idempotent and safe to call from any goroutine.
// Closing the connection makes the read loop return, which unregisters the client.
func (c *Client) Close() {
	c.closeOnce.Do(
		func() {
			close(c.chDone)

			_ = c.conn.Close()
		},
	)
}

// Send queues msg for delivery without blocking.
// If the queue is full the client is too slow and is disconnected.
// msg must not be modified after the call.
func (c *Client) Send(msg []byte) bool {
	select {
	case <-c.chDone:
		return false
	default:
	}

	select {
	case c.chSend <- msg:
		return true
	default:
		c.Close()

		return false
	}
}

// SendString is a convenience wrapper; it allocates once for the conversion.
func (c *Client) SendString(msg string) bool {
	return c.Send([]byte(msg))
}

func (c *Client) write(messageType int, payload []byte) bool {
	if errWriteDeadline := c.conn.SetWriteDeadline(time.Now().Add(_WaitWrite)); errWriteDeadline != nil {
		return false
	}

	return c.conn.WriteMessage(messageType, payload) == nil
}

func (c *Client) writeLoop() {
	ticker := time.NewTicker(_WaitPing)

	defer func() {
		ticker.Stop()
		c.Close()
	}()

	for {
		select {
		case msg := <-c.chSend:
			if !c.write(websocket.TextMessage, msg) {
				return
			}

		case <-ticker.C:
			if !c.write(websocket.PingMessage, nil) {
				return
			}

		case <-c.chDone:
			return
		}
	}
}
