package svc

import (
	"regexp"
)

func ValidateUserInput(u *UserAPIRequestType) []string {
	var errs []string
	if u.FirstName == "" {
		errs = append(errs, FirstNameRequiredError.Error())
	}
	if u.LastName == "" {
		errs = append(errs, LastNameRequiredError.Error())
	}
	if u.Email == "" {
		errs = append(errs, EmailRequiredError.Error())
	}
	if u.Password == "" {
		errs = append(errs, PasswordRequiredError.Error())
	}
	// email format
	if match, _ := regexp.MatchString(`[a-z0-9]+@[a-z]+\.[a-z]{2,3}`, u.Email); match {
		errs = append(errs, InvalidEmailFormatError.Error())
	}

	return errs
}

func ValidateCreateItemInput(u *ItemServiceRequestType) []string {
	var errs []string
	if u.Name == "" {
		errs = append(errs, FirstNameRequiredError.Error())
	}

	return errs
}
