package usecase

import (
	"go-ddd-ws-template/src/domain/repository"
	"log"

	"github.com/labstack/echo"
)

type ConnectionUsecase interface {
	HandleConnection(c echo.Context) error
}

type connectionUsecase struct {
	connectionRepo repository.ConnectionRepository
}

func NewConnectionUsecase(connectionRepo repository.ConnectionRepository) ConnectionUsecase {
	connectionUsecase := connectionUsecase{connectionRepo: connectionRepo}
	return &connectionUsecase
}
func (u *connectionUsecase) HandleConnection(c echo.Context) error {
	conn, err := u.connectionRepo.UpgradeProtocol(c)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err) // アップグレード失敗時にエラーログを記録
		return err                                          // エラーを返して処理を中断
	}
	defer conn.Close()

	// クライアントを追加
	u.connectionRepo.AddClient(conn)
	log.Println("New clent connected")

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
