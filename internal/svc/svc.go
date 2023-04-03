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

func (s *svc) CreateUser(ctx context.Context, user *UserServiceRequestType) (*UserServiceRequestType, error) {
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

func (s *svc) FindUserByEmail(ctx context.Context, email string) (*UserServiceRequestType, error) {
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

func (s *svc) hashPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		s.log.WithError(err).Error("failed to hash password")
		return "", err
	}

	return string(hash), nil
}
