package repository

import "github.com/gorilla/websocket"

type MessageRepository interface {
	AddClient(conn *websocket.Conn)
	RemoveClient(conn *websocket.Conn)
	BroadcastMessage(sender *websocket.Conn, message string)
}
