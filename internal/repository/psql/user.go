package psql

import (
	"context"
	"database/sql"
	"time"

	repo "github.com/danielwangai/todo-app/internal/repository"
)

func (dao *dbClient) CreateUser(ctx context.Context, user *repo.UserModelType) (*repo.UserModelType, error) {
	dao.log.Infof("Create User OBJ: %v", user)
	// check if email is unique
	x, err := dao.findUserByEmail(ctx, user.Email)
	if err != sql.ErrNoRows && err != nil {
		dao.log.Infof("Find USER Error: %v", err)
		return nil, err
	}
	dao.log.Infof("Find USER: %v", x)
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

	dao.log.Infof("DB-OP: User created successfully")
	return user, nil
}

func (dao *dbClient) findUserByEmail(ctx context.Context, email string) (*repo.UserModelType, error) {
	var id int
	var firstName, lastName string
	var createdAt time.Time
	query := `SELECT id, first_name, last_name, created_at FROM users WHERE email=$1`
	row := dao.db.QueryRow(query, email)
	err := row.Scan(&id, &firstName, &lastName, &createdAt)
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
		CreatedAt: createdAt,
	}
	dao.log.Info("-=----====: ", user)

	return user, nil
}

// `
// CREATE TABLE users(
// 	id SERIAL PRIMARY KEY,
// 	first_name VARCHAR(50) NOT NULL,
// 	last_name VARCHAR(50) NOT NULL,
// 	email VARCHAR(150) UNIQUE NOT NULL,
// 	password VARCHAR(200) NOT NULL,
// 	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
// 	updated_at TIMESTAMP
// )
// `
