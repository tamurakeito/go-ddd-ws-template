package repository

import (
	"github.com/labstack/echo"
)

type ConnectionRepository interface {
	HandleConnections(c echo.Context) error
}
