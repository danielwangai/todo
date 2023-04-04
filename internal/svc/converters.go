package svc

import (
	repo "github.com/danielwangai/todo-app/internal/repository"
)

func ConvertUserServiceToUserModelObject(userSvcObj *UserServiceType) *repo.UserModelType {
	return &repo.UserModelType{
		FirstName: userSvcObj.FirstName,
		LastName:  userSvcObj.LastName,
		Email:     userSvcObj.Email,
		Password:  userSvcObj.Password,
	}
}

func ConvertUserModelToUserServiceObject(userModel *repo.UserModelType) *UserServiceType {
	return &UserServiceType{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
		Password:  userModel.Password,
		CreatedAt: userModel.CreatedAt,
	}
}

// item conversions
func ConvertItemServiceRequestToModelObject(itemSvcObj *ItemServiceRequestType) *repo.ItemModelType {
	return &repo.ItemModelType{
		Name:      itemSvcObj.Name,
		UserId:    itemSvcObj.UserId,
		CreatedAt: itemSvcObj.CreatedAt,
	}
}

func ConvertItemServiceResponseToModelObject(itemSvcObj *ItemServiceResponseType) *repo.ItemModelType {
	return &repo.ItemModelType{
		ID:        itemSvcObj.ID,
		Name:      itemSvcObj.Name,
		UserId:    itemSvcObj.UserId,
		IsDeleted: itemSvcObj.IsDeleted,
		CreatedAt: itemSvcObj.CreatedAt,
		// UpdatedAt: itemSvcObj.UpdatedAt,
	}
}

func ConvertItemModelToServiceResponseObject(itemModel *repo.ItemModelType) *ItemServiceResponseType {
	item := &ItemServiceResponseType{
		ID:        itemModel.ID,
		Name:      itemModel.Name,
		UserId:    itemModel.UserId,
		IsDeleted: itemModel.IsDeleted,
		CreatedAt: itemModel.CreatedAt,
	}
	// if itemModel.UpdatedAt.Valid {
	// 	item.UpatedAt = itemModel.UpdatedAt.Time
	// }

	return item
}
