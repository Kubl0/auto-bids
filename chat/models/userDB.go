package models

type UserDB struct {
	Id    string      `json:"_id" bson:"_id"`
	Email string      `json:"email" bson:"email"`
	Rooms []UserRooms `json:"rooms" bson:"rooms"`
}

type PostUserDB struct {
	Email string      `json:"email" bson:"email" validate:"required,email"`
	Rooms []UserRooms `json:"rooms" bson:"rooms"`
}

type UserRooms struct {
	Id    string `json:"id" bson:"id" validate:"required"`
	Email string `json:"email" bson:"email" validate:"required,email"`
}
