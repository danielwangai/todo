package svc

import "time"

type UserAPIRequestType struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserLoginAPIRequestType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserServiceRequestType struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type UserServiceResponseType struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
}

// todo item types
type ItemServiceRequestType struct {
	Name      string
	UserId    int
	CreatedAt time.Time
}

type ItemServiceResponseType struct {
	ID        int
	Name      string
	UserId    int
	IsDeleted bool
	CreatedAt time.Time
	UpatedAt  time.Time
}
