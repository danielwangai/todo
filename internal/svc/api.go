package svc

import (
	"context"
)

type Svc interface {
	CreateUser(ctx context.Context, user *UserServiceRequestType) (*UserServiceRequestType, error)
	FindUserByEmail(ctx context.Context, email string) (*UserServiceRequestType, error)
	CreateTodoItem(ctx context.Context, item *ItemServiceRequestType) (*ItemServiceResponseType, error)
	GetAllTodoItems(ctx context.Context, id int) ([]*ItemServiceResponseType, error)
	FindTodoItemById(ctx context.Context, id int) (*ItemServiceResponseType, error)
	DeleteTodoItemById(ctx context.Context, id int) error
}
