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

// // item conversions
// func ConvertItemServiceRequestToModelObject(itemSvcObj *svc.ItemServiceRequestType) *svc.ItemServiceRequestType {
// 	return &repo.ItemModelType{
// 		Name:      itemSvcObj.Name,
// 		UserId:    itemSvcObj.UserId,
// 		CreatedAt: itemSvcObj.CreatedAt,
// 	}
// }
