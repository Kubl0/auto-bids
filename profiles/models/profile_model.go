package models

type Profile struct {
	Id           string `json:"_id" bson:"_id" validate:"required"`
	UserName     string `json:"user_name" bson:"user_name" validate:"required"`
	Email        string `json:"email" bson:"email" validate:"required,email"`
	ProfileImage string `json:"profile_image" bson:"profile_image" validate:"required"`
}
