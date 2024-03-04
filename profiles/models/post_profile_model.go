package models

type PostProfile struct {
	UserName     string `json:"user_name" bson:"user_name" validate:"required, max=30"`
	Email        string `json:"email" bson:"email" validate:"required,email"`
	ProfileImage string `json:"profile_image" bson:"profile_image" validate:"required, uri"`
}
