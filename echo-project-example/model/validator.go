package model

import (
	"gopkg.in/go-playground/validator.v9"
)

// Validator is
type Validator struct {
	validator *validator.Validate
}

// Validate is
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// NewValidator returns new validator instance
func NewValidator() *Validator {
	return &Validator{validator: validator.New()}
}
