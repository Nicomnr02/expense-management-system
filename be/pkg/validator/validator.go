package validator

import (
	"errors"
	"fmt"

	vd "github.com/go-playground/validator/v10"
)

type Validator interface {
	ValidateStruct(s interface{}) error
}
type validator struct {
	Validate *vd.Validate
}

func (v *validator) ValidateStruct(s interface{}) error {
	err := v.Validate.Struct(s)
	if err == nil {
		return nil
	}

	var msg string

	validationErrors, ok := err.(vd.ValidationErrors)
	if !ok || len(validationErrors) == 0 {
		return err
	}

	e := validationErrors[0]

	switch e.Tag() {
	case "required":
		msg = fmt.Sprintf("%s is required", e.Field())
	case "gt":
		msg = fmt.Sprintf("%s must be greater than %s", e.Field(), e.Param())
	case "min":
		msg = fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
	case "max":
		msg = fmt.Sprintf("%s cannot be more than %s characters", e.Field(), e.Param())
	case "url":
		msg = fmt.Sprintf("%s must be a valid URL", e.Field())
	default:
		msg = fmt.Sprintf("%s is invalid", e.Field())
	}

	return errors.New(msg)
}

func New() Validator {
	return &validator{
		Validate: vd.New(),
	}
}
