package svc

import "time"

type UserAPIRequestType struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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
