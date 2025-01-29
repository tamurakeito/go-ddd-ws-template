package usecase

import (
	"go-ddd-ws-template/src/core"
	"go-ddd-ws-template/src/repository"
	"log"

	"github.com/gorilla/websocket"
)

type ConnectionUsecase interface {
	HandleConnection(server *core.Server, conn *websocket.Conn) error
}

type connectionUsecase struct {
	connectionRepo repository.ConnectionRepository
}

func NewConnectionUsecase(connectionRepo repository.ConnectionRepository) ConnectionUsecase {
	connectionUsecase := connectionUsecase{connectionRepo: connectionRepo}
	return &connectionUsecase
}

func (u *connectionUsecase) HandleConnection(server *core.Server, conn *websocket.Conn) error {
	// クライアントを追加
	u.connectionRepo.AddClient(server, conn)
	log.Println("New client connected")

	// クライアントからのメッセージを受信してブロードキャスト
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		log.Printf("Received message: %s", message)
		server.BroadcastMessage(conn, string(message))
	}

	// クライアントを削除
	// server.Mutex.Lock()
	// delete(server.Clients, conn)
	// server.Mutex.Unlock()
	u.connectionRepo.RemoveClient(server, conn)
	log.Println("Client disconnected")
	return nil
}
