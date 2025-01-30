package usecase

import (
	"errors"
)

type UsecaseErr error

var (
	ErrUpgradeProtocol UsecaseErr = errors.New("ErrUpgradeProtocol")
	ErrHandleMessage   UsecaseErr = errors.New("ErrHandleMessage")
)
