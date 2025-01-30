package main

import (
	"fmt"
	"go-ddd-ws-template/src/injector"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("server start")
	httpHandler := injector.InjectHttpHandler()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ws", httpHandler.HandleConnection())

	port := ":8081"
	log.Printf("WebSocket server is running on http://localhost%s/ws", port)
	e.Logger.Fatal(e.Start(port))
}
