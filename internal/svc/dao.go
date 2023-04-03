package svc

import (
	"context"

	repo "github.com/danielwangai/todo-app/internal/repository"
)

type DAO interface {
	CreateUser(ctx context.Context, user *repo.UserModelType) (*repo.UserModelType, error)
	FindUserByEmail(ctx context.Context, email string) (*repo.UserModelType, error)
	CreateTodoItem(ctx context.Context, item *repo.ItemModelType) (*repo.ItemModelType, error)
	GetAllTodoItems(ctx context.Context, id int) ([]*repo.ItemModelType, error)
	FindTodoItemById(ctx context.Context, id int) (*repo.ItemModelType, error)
	DeleteTodoItemById(ctx context.Context, id int) error
}
