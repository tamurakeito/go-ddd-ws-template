package infrastructure

import (
	"go-ddd-ws-template/src/domain/entity"
	"go-ddd-ws-template/src/domain/repository"
	"go-ddd-ws-template/src/infrastructure"
	"log"

	"github.com/labstack/echo"
)

type ConnectionRepository struct {
	server *infrastructure.Server
}

func NewConnectionRepository(server *infrastructure.Server) repository.ConnectionRepository {
	connectionRepository := ConnectionRepository{server: server}
	return &connectionRepository
}

func (repo *ConnectionRepository) UpgradeProtocol(c echo.Context) (client entity.Client, err error) {
	conn, err := repo.server.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	client = infrastructure.NewClient(conn)
	return
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

func (repo *ConnectionRepository) HandleMessage(client entity.Client) (err error) {
	message, err := client.ReadMessage()
	if err != nil {
		log.Printf("Error reading message: %v", err)
		return
	}
	log.Printf("Received message: %s", message)
	repo.server.BroadcastMessage(client, string(message))
	return
}
