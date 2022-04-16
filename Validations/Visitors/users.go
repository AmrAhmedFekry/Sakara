package Visitors

import (
	"GoGin/Models"
	"GoGin/Validations"

	validation "github.com/go-ozzo/ozzo-validation"
)

func RegisterValidation(user Models.User) validation.Errors {
	return validation.Errors{
		"username": validation.Validate(user.UserName, Validations.RequiredRule(), Validations.MinMaxRule(), Validations.Email()),
		"password": validation.Validate(user.Password, Validations.RequiredRule(), Validations.MinMaxRule()),
		"deposit":  validation.Validate(user.Deposit, Validations.RequiredRule()),
	}
}

func LoginValidation(user Models.User) validation.Errors {
	return validation.Errors{
		"username": validation.Validate(user.UserName, Validations.RequiredRule(), Validations.MinMaxRule(), Validations.Email()),
		"password": validation.Validate(user.Password, Validations.RequiredRule(), Validations.MinMaxRule()),
	}
}
