package handler

import (
	"errors"
	"go-ddd-ws-template/src/presentation/api_error"
	"go-ddd-ws-template/src/usecase"
	"net/http"

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
		err := handler.connectionUsecase.HandleConnection(c)
		if err != nil {
			if errors.Is(err, usecase.ErrUpgradeProtocol) {
				return c.JSON(api_error.NewWebSocketHandshakeError())
			} else if errors.Is(err, usecase.ErrHandleMessage) {
				return c.JSON(api_error.NewMessageHandlingError())
			}
		}
		return c.JSON(http.StatusOK, nil)
	}
}
