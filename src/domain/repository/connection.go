package repository

import (
	"go-ddd-ws-template/src/domain/entity"

	"github.com/labstack/echo"
)

type ConnectionRepository interface {
	UpgradeProtocol(c echo.Context) (client entity.ClientInterface, err error)
	AddClient(client entity.ClientInterface)
	RemoveClient(client entity.ClientInterface)
	HandleMessage(client entity.ClientInterface) (err error)
}
