package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	Validate *validator.Validate
}

func New() Validator {
	return Validator{
		Validate: validator.New(),
	}
}
