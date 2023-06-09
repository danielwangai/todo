package svc

import (
	"context"
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type svc struct {
	dao DAO
	log *logrus.Logger
}

// New returns a new Svc object
func New(dao DAO, log *logrus.Logger) Svc {
	return &svc{dao, log}
}

func (s *svc) CreateUser(ctx context.Context, user *UserServiceType) (*UserServiceType, error) {
	// hash password
	hash, err := s.hashPassword([]byte(user.Password))
	if err != nil {
		return nil, err
	}

	user.Password = hash
	// convert service request to DB model
	usrModel := ConvertUserServiceToUserModelObject(user)

	newUser, err := s.dao.CreateUser(ctx, usrModel)
	if err != nil {
		return nil, err
	}

	// convert user model back to svc model
	user = ConvertUserModelToUserServiceObject(newUser)

	// return ConvertUserModelToUserServiceObject(newUser), nil
	return user, nil
}

func (s *svc) FindUserById(ctx context.Context, id int) (*UserServiceType, error) {
	user, err := s.dao.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	return ConvertUserModelToUserServiceObject(user), nil
}

func (s *svc) CreateTodoItem(ctx context.Context, item *ItemServiceRequestType) (*ItemServiceResponseType, error) {
	// convert service request to item model
	itemModel := ConvertItemServiceRequestToModelObject(item)
	itemModel.CreatedAt = time.Now()

	// create item
	itemModel, err := s.dao.CreateTodoItem(ctx, itemModel)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}

	return ConvertItemModelToServiceResponseObject(itemModel), nil
}

func (s *svc) FindUserByEmail(ctx context.Context, email string) (*UserServiceType, error) {
	user, err := s.dao.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return ConvertUserModelToUserServiceObject(user), nil
}

func (s *svc) GetAllTodoItems(ctx context.Context, id int) ([]*ItemServiceResponseType, error) {
	items, err := s.dao.GetAllTodoItems(ctx, id)
	if err != nil {
		return nil, err
	}

	itemList := []*ItemServiceResponseType{}
	for _, v := range items {
		itemList = append(itemList, ConvertItemModelToServiceResponseObject(v))
	}

	return itemList, nil
}

func (s *svc) FindTodoItemById(ctx context.Context, id int) (*ItemServiceResponseType, error) {
	item, err := s.dao.FindTodoItemById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			s.log.WithError(err).Error("item matching id: %d not found", id)
			return nil, err
		}

		s.log.WithError(err).Error("an error ocurred when fetching item matching id: %d", id)
		return nil, err
	}

	return ConvertItemModelToServiceResponseObject(item), err
}

func (s *svc) DeleteTodoItemById(ctx context.Context, id int) error {
	err := s.dao.DeleteTodoItemById(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *svc) UpdateTodoItem(ctx context.Context, item *ItemServiceResponseType, newName string) (*ItemServiceResponseType, error) {
	itemModel := ConvertItemServiceResponseToModelObject(item)
	itemModel, err := s.dao.UpdateTodoItem(ctx, itemModel, newName)
	if err != nil {
		s.log.WithError(err).Error("an error ocurred when updating item matching id: %d", item.ID)
		return nil, err
	}

	return ConvertItemModelToServiceResponseObject(itemModel), err
}

func (s *svc) hashPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		s.log.WithError(err).Error("failed to hash password")
		return "", err
	}

	return string(hash), nil
}
