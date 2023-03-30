package svc

import (
	"context"

	repo "github.com/danielwangai/todo-app/internal/repository"
)

type DAO interface {
	CreateUser(ctx context.Context, user *repo.UserModelType) (*repo.UserModelType, error)
}