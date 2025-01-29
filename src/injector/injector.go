package injector

import (
	"go-ddd-ws-template/src/infrastructure"
	"go-ddd-ws-template/src/presentation/handler"
	"go-ddd-ws-template/src/repository"
	"go-ddd-ws-template/src/usecase"
)

func InjectConnectionRepository() repository.ConnectionRepository {
	return infrastructure.NewConnectionRepository()
}

func InjectConnectionUsecase() usecase.ConnectionUsecase {
	connectionRepo := InjectConnectionRepository()
	return usecase.NewConnectionUsecase(connectionRepo)
}

func InjectHttpHandler() handler.HttpHandler {
	return handler.NewHttpHandler(InjectConnectionUsecase())
}
