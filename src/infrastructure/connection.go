package infrastructure

import (
	"go-ddd-ws-template/src/core"
	"go-ddd-ws-template/src/repository"
	"log"

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

func (repo *ConnectionRepository) HandleMessage(server *core.Server, conn *websocket.Conn) (err error) {
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Printf("Error reading message: %v", err)
		return
	}
	log.Printf("Received message: %s", message)
	server.BroadcastMessage(conn, string(message))
	return
}
