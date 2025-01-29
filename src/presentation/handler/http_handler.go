package handler

import (
	"go-ddd-ws-template/src/infrastructure"
	"go-ddd-ws-template/src/usecase"
	"log"

	"github.com/labstack/echo"
)

type HttpHandler struct {
	connectionUsecase usecase.ConnectionUsecase
}

func NewHttpHandler(connectionUsecase usecase.ConnectionUsecase) HttpHandler {
	httpHandler := HttpHandler{connectionUsecase: connectionUsecase}
	return httpHandler
}

func (handler *HttpHandler) HandleConnection() echo.HandlerFunc {
	return func(c echo.Context) error {
		// 新しいWebSocketサーバーインスタンスを作成
		server := infrastructure.NewServer()

		// HTTP接続をWebSocket接続にアップグレード
		conn, err := server.Upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			log.Printf("Error upgrading to WebSocket: %v", err) // アップグレード失敗時にエラーログを記録
			return err                                          // エラーを返して処理を中断
		}
		defer conn.Close() // 関数の終了時にWebSocket接続を確実に閉じる

		err = handler.connectionUsecase.HandleConnection(server, conn)

		return err
	}
}
