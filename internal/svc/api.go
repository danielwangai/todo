package svc

import "context"

type Svc interface {
	CreateUser(ctx context.Context, user *UserServiceRequestType) (*UserServiceRequestType, error)
	CreateTodoItem(ctx context.Context, item *ItemServiceRequestType) (*ItemServiceResponseType, error)
}
