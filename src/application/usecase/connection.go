package usecase

import (
	"go-ddd-ws-template/domain/repository"

	"github.com/labstack/echo"
)

type ConnectionUsecase interface {
	Connetion(c echo.Context) error
}

type connectionUsecase struct {
	connectionRepo repository.ConnectionRepository
}

func NewConnectionUsecase(connectionRepo repository.ConnectionRepository) ConnectionUsecase {
	connectionUsecase := connectionUsecase{connectionRepo: connectionRepo}
	return &connectionUsecase
}

func (u *connectionUsecase) Connetion(c echo.Context) error {
	return u.connectionRepo.HandleConnections(c)
}
