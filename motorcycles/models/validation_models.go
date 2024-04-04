package models

type Id struct {
	Id string `json:"id" bson:"_id" validation:"required"`
}

type Email struct {
	Email string `validate:"required,email"`
}
