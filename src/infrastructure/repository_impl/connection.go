package infrastructure

import (
	"errors"
	"go-ddd-ws-template/src/domain"
	"go-ddd-ws-template/src/domain/entity"
	"go-ddd-ws-template/src/domain/repository"
	"go-ddd-ws-template/src/infrastructure"
	"io"
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

func (repo *ConnectionRepository) UpgradeProtocol(c echo.Context) (entity.Client, error) {
	conn, err := repo.server.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return nil, domain.ErrInternal
	}
	client := infrastructure.NewClient(conn)
	return client, nil
}

func (repo *ConnectionRepository) AddClient(client entity.Client) {
	repo.server.Mutex.Lock()
	repo.server.Clients[client] = true
	repo.server.Mutex.Unlock()
}

func (repo *ConnectionRepository) RemoveClient(client entity.Client) {
	repo.server.Mutex.Lock()
	delete(repo.server.Clients, client)
	repo.server.Mutex.Unlock()
}

func (repo *ConnectionRepository) HandleMessage(client entity.Client) error {
	message, err := client.ReadMessage()
	if err != nil {
		if errors.Is(err, io.EOF) || websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
			log.Println("Client disconnected gracefully")
			return domain.EOF
		}
		log.Printf("Error reading message: %v", err)
		return domain.ErrConnection
	}
	log.Printf("Received message: %s", message)
	repo.server.BroadcastMessage(client, string(message))
	return nil
}
