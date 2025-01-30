package domain

import (
	"errors"
)

type RepositoryErr error

var (
	EOF RepositoryErr = errors.New("EOF")

	ErrConnection RepositoryErr = errors.New("ErrConnection")

	// // 指定されたリソースがない
	// ErrResourceNotFound RepositoryErr = errors.New("ErrResourceNotFound")

	// // リソースの競合が発生した場合（データベースのユニーク制約違反）
	// ErrResourceConflict RepositoryErr = errors.New("ErrResourceConflict")

	// // データベースが利用できない
	// ErrDatabaseUnavailable RepositoryErr = errors.New("ErrDatabaseUnavailable")

	// システム内部で想定外のエラー
	ErrInternal RepositoryErr = errors.New("ErrInternal")
)
