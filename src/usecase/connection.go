package usecase

import (
	"go-ddd-ws-template/src/domain/repository"
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type ConnectionUsecase interface {
	UpgradeProtocol(c echo.Context) (conn *websocket.Conn, err error)
	ManageClient(conn *websocket.Conn) error
}

type connectionUsecase struct {
	connectionRepo repository.ConnectionRepository
}

func NewConnectionUsecase(connectionRepo repository.ConnectionRepository) ConnectionUsecase {
	connectionUsecase := connectionUsecase{connectionRepo: connectionRepo}
	return &connectionUsecase
}
func (u *connectionUsecase) UpgradeProtocol(c echo.Context) (conn *websocket.Conn, err error) {
	conn, err = u.connectionRepo.UpgradeProtocol(c)
	return
}
func (u *connectionUsecase) ManageClient(conn *websocket.Conn) error {
	// クライアントを追加
	u.connectionRepo.AddClient(conn)
	log.Println("New client connected")

	// クライアントからのメッセージを受信してブロードキャスト
	for {
		err := u.connectionRepo.HandleMessage(conn)
		if err != nil {
			break
		}
	}

	// クライアントを削除
	u.connectionRepo.RemoveClient(conn)
	log.Println("Client disconnected")
	return nil
}
