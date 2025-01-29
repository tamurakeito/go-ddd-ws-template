package infrastructure

import (
	"go-ddd-ws-template/src/domain/repository"
	"go-ddd-ws-template/src/infrastructure"
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type ConnectionRepository struct {
	server *infrastructure.Server
}

func NewConnectionRepository(server *infrastructure.Server) repository.ConnectionRepository {
	connectionRepository := ConnectionRepository{server: server}
	return &connectionRepository
}

// HTTP接続をWebSocket接続にアップグレード
func (repo *ConnectionRepository) UpgradeProtocol(c echo.Context) (conn *websocket.Conn, err error) {
	conn, err = repo.server.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err) // アップグレード失敗時にエラーログを記録
		return                                              // エラーを返して処理を中断
	}
	return
}

func (repo *ConnectionRepository) AddClient(conn *websocket.Conn) {
	repo.server.Mutex.Lock()
	repo.server.Clients[conn] = true
	repo.server.Mutex.Unlock()
}

func (repo *ConnectionRepository) RemoveClient(conn *websocket.Conn) {
	repo.server.Mutex.Lock()
	delete(repo.server.Clients, conn)
	repo.server.Mutex.Unlock()
}

func (repo *ConnectionRepository) HandleMessage(conn *websocket.Conn) (err error) {
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Printf("Error reading message: %v", err)
		return
	}
	log.Printf("Received message: %s", message)
	repo.server.BroadcastMessage(conn, string(message))
	return
}
