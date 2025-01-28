package infrastructure

import (
	"go-ddd-ws-template/src/domain/repository"
	"sync"

	"github.com/gorilla/websocket"
)

type MessageRepository struct {
	Clients map[*websocket.Conn]bool // 接続中のクライアントを保持するマップ
	Mutex   sync.Mutex               // Clientsへの安全なアクセスを保証するためのミューテックス
}

func NewMessagingRepository() repository.MessageRepository {
	messageRepository := MessageRepository{}
	return &messageRepository
}

func (repo *MessageRepository) AddClient(conn *websocket.Conn) {
	repo.Mutex.Lock()
	repo.Clients[conn] = true
	repo.Mutex.Unlock()
}

func (repo *MessageRepository) RemoveClient(conn *websocket.Conn) {
	repo.Mutex.Lock()
	delete(repo.Clients, conn)
	repo.Mutex.Unlock()
}

func (repo *MessageRepository) BroadcastMessage(sender *websocket.Conn, message string) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	for client := range repo.Clients {
		if client != sender {
			client.WriteMessage(websocket.TextMessage, []byte(message))
		}
	}
}
