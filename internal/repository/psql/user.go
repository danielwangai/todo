package psql

import (
	"context"
	"database/sql"
	"time"

	repo "github.com/danielwangai/todo-app/internal/repository"
)

func (dao *dbClient) CreateUser(ctx context.Context, user *repo.UserModelType) (*repo.UserModelType, error) {
	// check if email is unique
	_, err := dao.FindUserByEmail(ctx, user.Email)
	if err != sql.ErrNoRows && err != nil {
		return nil, err
	}
	var id int

	query := `
	INSERT INTO users (first_name, last_name, email, password, created_at)
	VALUES($1, $2, $3, $4, $5) RETURNING id`
	err = dao.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt).Scan(&id)
	if err != nil {
		dao.log.WithError(err).Error("an error occurred when inserting user to the database")
		return nil, err
	}

	user.ID = id

	dao.log.Infof("DB-OP: User created successfully: %v", user)
	return user, nil
}

func (dao *dbClient) FindUserByEmail(ctx context.Context, email string) (*repo.UserModelType, error) {
	var id int
	var firstName, lastName, password string
	var createdAt time.Time
	query := `SELECT id, first_name, last_name, password, created_at FROM users WHERE email=$1`
	row := dao.db.QueryRow(query, email)
	err := row.Scan(&id, &firstName, &lastName, &password, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			dao.log.WithError(err).Error("DB-OP: a user with a similar email exists")
			return nil, err
		}
		dao.log.WithError(err).Error("DB-OP: an error ocurred when searching for user by email")
		return nil, err
	}
	user := &repo.UserModelType{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
	}

	return user, nil
}

func (dao *dbClient) FindUserById(ctx context.Context, id int) (*repo.UserModelType, error) {
	var firstName, lastName, email, password string
	var createdAt time.Time
	query := `SELECT id, first_name, last_name, email, password, created_at FROM users WHERE id=$1`
	row := dao.db.QueryRow(query, id)
	err := row.Scan(&id, &firstName, &lastName, &email, &password, &createdAt)
	if err != nil {
		return nil, err
	}
	user := &repo.UserModelType{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
	}

	return user, nil
}
