package usecase

import (
	"go-ddd-ws-template/src/infrastructure"
	"go-ddd-ws-template/src/domain/repository"
	"log"

	"github.com/gorilla/websocket"
)

type ConnectionUsecase interface {
	HandleConnection(server *infrastructure.Server, conn *websocket.Conn) error
}

type connectionUsecase struct {
	connectionRepo repository.ConnectionRepository
}

func NewConnectionUsecase(connectionRepo repository.ConnectionRepository) ConnectionUsecase {
	connectionUsecase := connectionUsecase{connectionRepo: connectionRepo}
	return &connectionUsecase
}

func (u *connectionUsecase) HandleConnection(server *infrastructure.Server, conn *websocket.Conn) error {
	// クライアントを追加
	u.connectionRepo.AddClient(server, conn)
	log.Println("New client connected")

	// クライアントからのメッセージを受信してブロードキャスト
	for {
		err := u.connectionRepo.HandleMessage(server, conn)
		if err != nil {
			break
		}
	}

	// クライアントを削除
	u.connectionRepo.RemoveClient(server, conn)
	log.Println("Client disconnected")
	return nil
}
