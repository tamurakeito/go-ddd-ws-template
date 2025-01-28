package presentation

import (
	"go-ddd-ws-template/src/presentation/handler"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, connectionHandler handler.ConnectionHandler) {
	e.GET("/ws", connectionHandler.Connetion())
}
