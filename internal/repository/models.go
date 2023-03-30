package repository

import (
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
