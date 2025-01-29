package injector

import (
	"go-ddd-ws-template/src/domain/repository"
	"go-ddd-ws-template/src/infrastructure"
	repository_impl "go-ddd-ws-template/src/infrastructure/repository_impl"
	"go-ddd-ws-template/src/presentation/handler"
	"go-ddd-ws-template/src/usecase"
)

func InjectServer() infrastructure.Server {
	return *infrastructure.NewServer()
}

func InjectConnectionRepository() repository.ConnectionRepository {
	server := InjectServer()
	return repository_impl.NewConnectionRepository(&server)
}

func InjectConnectionUsecase() usecase.ConnectionUsecase {
	connectionRepo := InjectConnectionRepository()
	return usecase.NewConnectionUsecase(connectionRepo)
}

func InjectHttpHandler() handler.HttpHandler {
	return handler.NewHttpHandler(InjectConnectionUsecase())
}
