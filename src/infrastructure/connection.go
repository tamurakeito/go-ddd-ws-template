package infrastructure

import (
	"go-ddd-ws-template/src/core"
	"go-ddd-ws-template/src/repository"

	"github.com/gorilla/websocket"
)

type ConnectionRepository struct{}

func NewConnectionRepository() repository.ConnectionRepository {
	connectionRepository := ConnectionRepository{}
	return &connectionRepository
}

func (repo *ConnectionRepository) AddClient(server *core.Server, conn *websocket.Conn) {
	server.Mutex.Lock()
	server.Clients[conn] = true
	server.Mutex.Unlock()
}

func (repo *ConnectionRepository) RemoveClient(server *core.Server, conn *websocket.Conn) {
	server.Mutex.Lock()
	delete(server.Clients, conn)
	server.Mutex.Unlock()
}

// func (repo *ConnectionRepository) BroadcastMessage(server *core.Server, sender *websocket.Conn, message string) {
// 	server.Mutex.Lock()
// 	defer server.Mutex.Unlock()

// 	for client := range server.Clients {
// 		if client != sender {
// 			client.WriteMessage(websocket.TextMessage, []byte(message))
// 		}
// 	}
// }
