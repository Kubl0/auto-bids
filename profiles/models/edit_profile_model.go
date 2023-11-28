package models

type EditProfile struct {
	UserName     string `bson:"user_name"`
	Email        string `bson:"email"`
	ProfileImage string `bson:"profile_image"`
}
