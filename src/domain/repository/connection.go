package repository

import (
	"go-ddd-ws-template/src/infrastructure"

	"github.com/gorilla/websocket"
)

type ConnectionRepository interface {
	AddClient(server *infrastructure.Server, conn *websocket.Conn)
	RemoveClient(server *infrastructure.Server, conn *websocket.Conn)
	HandleMessage(server *infrastructure.Server, conn *websocket.Conn) (err error)
}
