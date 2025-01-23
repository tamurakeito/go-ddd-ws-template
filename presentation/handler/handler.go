package handler

import (
	"go-ddd-ws-template/application/usecase"

	"github.com/labstack/echo"
)

type ConnectionHandler struct {
	connectionUsecase usecase.ConnectionUsecase
}

func NewConnectionHandler(connectionUsecase usecase.ConnectionUsecase) ConnectionHandler {
	connectionHandler := ConnectionHandler{connectionUsecase: connectionUsecase}
	return connectionHandler
}

func (handler *ConnectionHandler) Connetion() echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler.connectionUsecase.Connetion(c)
	}
}
