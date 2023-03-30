package svc

import "context"

type Svc interface {
	CreateUser(ctx context.Context, user *UserServiceRequestType) (*UserServiceRequestType, error)
}
