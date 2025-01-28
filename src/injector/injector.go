package injector

import (
	"go-ddd-ws-template/src/application/usecase"
	"go-ddd-ws-template/src/domain/repository"
	"go-ddd-ws-template/src/infrastructure"
	"go-ddd-ws-template/src/presentation/handler"

	"github.com/gorilla/websocket"
)

func InjectServer() infrastructure.Server {
	repo := &infrastructure.MessageRepository{
		Clients: make(map[*websocket.Conn]bool),
	}
	return *infrastructure.NewServer(repo)
}

func InjectConnectionRepository() repository.ConnectionRepository {
	server := InjectServer()
	return infrastructure.NewConnectionRepository(&server)
}

func InjectConnectionUsecase() usecase.ConnectionUsecase {
	connectionRepo := InjectConnectionRepository()
	return usecase.NewConnectionUsecase(connectionRepo)
}

func InjectConnectionHandler() handler.ConnectionHandler {
	return handler.NewConnectionHandler(InjectConnectionUsecase())
}
