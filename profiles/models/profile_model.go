package models

type Profile struct {
	Id           string `bson:"_id" validate:"required"`
	UserName     string `bson:"user_name" validate:"required"`
	Email        string `bson:"email" validate:"required,email"`
	ProfileImage string `bson:"profile_image" validate:"required"`
}
