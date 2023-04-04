package svc

import (
	"context"
)

type Svc interface {
	CreateUser(ctx context.Context, user *UserServiceType) (*UserServiceType, error)
	FindUserByEmail(ctx context.Context, email string) (*UserServiceType, error)
	CreateTodoItem(ctx context.Context, item *ItemServiceRequestType) (*ItemServiceResponseType, error)
	GetAllTodoItems(ctx context.Context, id int) ([]*ItemServiceResponseType, error)
	FindTodoItemById(ctx context.Context, id int) (*ItemServiceResponseType, error)
	DeleteTodoItemById(ctx context.Context, id int) error
	UpdateTodoItem(ctx context.Context, item *ItemServiceResponseType, newName string) (*ItemServiceResponseType, error)
	FindUserById(ctx context.Context, id int) (*UserServiceType, error)
}
