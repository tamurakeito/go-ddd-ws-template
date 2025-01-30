package usecase

import (
	"go-ddd-ws-template/src/domain/repository"
	"log"

	"github.com/labstack/echo"
)

type ConnectionUsecase interface {
	HandleConnection(c echo.Context) (err error)
}

type connectionUsecase struct {
	connectionRepo repository.ConnectionRepository
}

func NewConnectionUsecase(connectionRepo repository.ConnectionRepository) ConnectionUsecase {
	connectionUsecase := connectionUsecase{connectionRepo: connectionRepo}
	return &connectionUsecase
}

func (u *connectionUsecase) HandleConnection(c echo.Context) error {
	client, err := u.connectionRepo.UpgradeProtocol(c)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err) // アップグレード失敗時にエラーログを記録
		return ErrUpgradeProtocol                           // エラーを返して処理を中断
	}
	defer client.Close()

	// クライアントを追加
	u.connectionRepo.AddClient(client)
	log.Println("New clent connected")

	// クライアントからのメッセージを受信してブロードキャスト
	for {
		err = u.connectionRepo.HandleMessage(client)
		if err != nil {
			break
		}
	}

	// クライアントを削除
	u.connectionRepo.RemoveClient(client)
	log.Println("Client disconnected")
	if err != nil {
		return ErrHandleMessage
	} else {
		return nil
	}
}
