package models

type Email struct {
	Email string `validate:"required,email"`
}
