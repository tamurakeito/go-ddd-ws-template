package repository

import (
	"go-ddd-ws-template/src/core"

	"github.com/gorilla/websocket"
)

type ConnectionRepository interface {
	AddClient(server *core.Server, conn *websocket.Conn)
	RemoveClient(server *core.Server, conn *websocket.Conn)
	// BroadcastMessage(server *core.Server, sender *websocket.Conn, message string)
}
