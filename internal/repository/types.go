package repository

import (
	"database/sql"
	"time"
)

// User struct contains properties of a user
type UserModelType struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ItemModelType struct {
	ID        int
	Name      string
	UserId    int
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	// UpdatedAt *time.Time
}
