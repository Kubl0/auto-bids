package models

type Profile struct {
	Id           string `bson:"_id"`
	UserName     string `bson:"user_name"`
	Email        string `bson:"email,email"`
	ProfileImage string `bson:"profile_image"`
}
