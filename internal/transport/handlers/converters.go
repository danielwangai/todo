package handlers

import (
	"github.com/danielwangai/todo-app/internal/svc"
)

func ConvertUserRequestToServiceObject(usrRequest *svc.UserAPIRequestType) *svc.UserServiceRequestType {
	return &svc.UserServiceRequestType{
		FirstName: usrRequest.FirstName,
		LastName:  usrRequest.LastName,
		Email:     usrRequest.Email,
		Password:  usrRequest.Password,
	}
}
