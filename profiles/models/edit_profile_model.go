package models

type EditProfile struct {
	UserName     string `bson:"user_name" validate:"required"`
	Email        string `bson:"email" validate:"required,email"`
	ProfileImage string `bson:"profile_image" validate:"required"`
}
