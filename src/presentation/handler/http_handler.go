package handler

import (
	"go-ddd-ws-template/src/core"
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

func (handler *HttpHandler) LinkConnection() echo.HandlerFunc {
	return func(c echo.Context) error {
		server := core.NewServer()

		// return server.HandleConnections(c)
		conn, err := server.Upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			log.Printf("Error upgrading to WebSocket: %v", err)
			return err
		}
		defer conn.Close()

		err = handler.connectionUsecase.HandleConnection(server, conn)

		return err
	}
}
