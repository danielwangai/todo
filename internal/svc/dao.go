package svc

import (
	"context"

	repo "github.com/danielwangai/todo-app/internal/repository"
)

type DAO interface {
	CreateUser(ctx context.Context, user *repo.UserModelType) (*repo.UserModelType, error)
	CreateTodoItem(ctx context.Context, item *repo.ItemModelType) (*repo.ItemModelType, error)
	FindUserByEmail(ctx context.Context, email string) (*repo.UserModelType, error)
}
