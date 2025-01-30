package repository

import (
	"go-ddd-ws-template/src/domain/entity"

	"github.com/labstack/echo"
)

type ConnectionRepository interface {
	UpgradeProtocol(c echo.Context) (client entity.Client, err error)
	AddClient(client entity.Client)
	RemoveClient(client entity.Client)
	HandleMessage(client entity.Client) (err error)
}
