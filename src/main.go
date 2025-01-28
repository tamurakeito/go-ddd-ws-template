package main

import (
	"fmt"
	"go-ddd-ws-template/injector"
	"go-ddd-ws-template/presentation"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("server start")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// WebSocketサーバーを初期化
	connectionHandler := injector.InjectConnectionHandler()

	// WebSocketのエンドポイントを設定
	presentation.InitRouting(e, connectionHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
