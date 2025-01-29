package handler

import (
	"go-ddd-ws-template/src/usecase"

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
		conn, err := handler.connectionUsecase.UpgradeProtocol(c)
		defer conn.Close() // 関数の終了時にWebSocket接続を確実に閉じる

		err = handler.connectionUsecase.ManageClient(conn)

		return err
	}
}
