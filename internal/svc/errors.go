package svc

import "errors"

var (
	FirstNameRequiredError  = errors.New("first name is required")
	LastNameRequiredError   = errors.New("last name is required")
	EmailRequiredError      = errors.New("email is required")
	PasswordRequiredError   = errors.New("password field is required")
	InvalidEmailFormatError = errors.New("invalid email format")
)
