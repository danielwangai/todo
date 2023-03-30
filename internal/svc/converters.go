package svc

import (
	repo "github.com/danielwangai/todo-app/internal/repository"
)

func ConvertUserServiceToUserModelObject(userSvcObj *UserServiceRequestType) *repo.UserModelType {
	return &repo.UserModelType{
		FirstName: userSvcObj.FirstName,
		LastName:  userSvcObj.LastName,
		Email:     userSvcObj.Email,
		Password:  userSvcObj.Password,
	}
}

func ConvertUserModelToUserServiceObject(userModel *repo.UserModelType) *UserServiceRequestType {
	return &UserServiceRequestType{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
		Password:  userModel.Password,
		CreatedAt: userModel.CreatedAt,
	}
}
