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
	s.log.Infof("Password: %v", user.Password)
	hash, err := s.hashPassword([]byte(user.Password))
	if err != nil {
		return nil, err
	}

	user.Password = hash
	s.log.Infof("Password Hashed: %v", user)
	// convert service request to DB model
	usrModel := ConvertUserServiceToUserModelObject(user)
	s.log.Infof("Converted to model obj: %v", usrModel)

	newUser, err := s.dao.CreateUser(ctx, usrModel)
	if err != nil {
		return nil, err
	}

	// convert user model back to svc model
	user = ConvertUserModelToUserServiceObject(newUser)
	s.log.Info("Converted USer ID 1: ", newUser)

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
