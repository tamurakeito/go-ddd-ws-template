package infrastructure

import (
	"fmt"
	"go-ddd-ws-template/src/domain/entity"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) entity.ClientInterface {
	return &Client{conn: conn}
}

func (c *Client) ReadMessage() (message []byte, err error) {
	_, message, err = c.conn.ReadMessage()
	return
}

func (c *Client) SendMessage(message string) error {
	if c.conn == nil {
		return fmt.Errorf("connection is nil")
	}
	return c.conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func (c *Client) Close() error {
	if c.conn == nil {
		return fmt.Errorf("connection is already closed")
	}
	return c.conn.Close()
}
