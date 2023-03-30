package svc

import (
	"context"

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

func (s *svc) hashPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		s.log.WithError(err).Error("failed to hash password")
		return "", err
	}

	return string(hash), nil
}
