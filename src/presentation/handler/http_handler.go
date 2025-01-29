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
		return handler.connectionUsecase.HandleConnection(c)
	}
}
