package repository

import (
	"go-ddd-ws-template/src/core"

	"github.com/gorilla/websocket"
)

type ConnectionRepository interface {
	AddClient(server *core.Server, conn *websocket.Conn)
	RemoveClient(server *core.Server, conn *websocket.Conn)
	HandleMessage(server *core.Server, conn *websocket.Conn) (err error)
}
