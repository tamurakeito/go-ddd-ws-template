package infrastructure

import (
	"go-ddd-ws-template/domain/repository"
	"log"

	"github.com/labstack/echo"
)

type ConnectionRepository struct {
	*Server
}

func NewConnectionRepository(server *Server) repository.ConnectionRepository {
	connectionRepository := ConnectionRepository{server}
	return &connectionRepository
}

func (repo *ConnectionRepository) HandleConnections(c echo.Context) error {

	// HTTP接続をWebSocketにアップグレード
	conn, err := repo.Server.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		return err
	}
	defer conn.Close() // 接続が終了したらクリーンアップ

	// 新しいクライアントをマップに追加
	repo.Server.repository.AddClient(conn)
	defer repo.Server.repository.RemoveClient(conn)
	// 	repo.Server.Mutex.Lock()
	// 	repo.Server.Clients[conn] = true
	// 	repo.Server.Mutex.Unlock()
	log.Println("New client connected")

	// クライアントからのメッセージを受信し続けるループ
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break // エラーが発生した場合はループを終了
		}
		log.Printf("Received message: %s", message)
		// メッセージを他のクライアントにブロードキャスト
		repo.Server.repository.BroadcastMessage(conn, string(message))
	}

	log.Println("Client disconnected")
	return nil
}
