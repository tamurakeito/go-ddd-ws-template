package repository

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type ConnectionRepository interface {
	UpgradeProtocol(c echo.Context) (conn *websocket.Conn, err error)
	AddClient(conn *websocket.Conn)
	RemoveClient(conn *websocket.Conn)
	HandleMessage(conn *websocket.Conn) (err error)
}
