package models

type EditProfile struct {
	UserName     string `bson:"user_name" validate:"required"`
	ProfileImage string `bson:"profile_image" validate:"required"`
}
