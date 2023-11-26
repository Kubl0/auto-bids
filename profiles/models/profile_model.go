package models

type Profile struct {
	Id           string `bson:"id"`
	UserName     string `bson:"user_name"`
	Email        string `bson:"email"`
	ProfileImage string `bson:"profile_image"`
}
