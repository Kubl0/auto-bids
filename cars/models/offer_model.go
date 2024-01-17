package models

type Offer struct {
	Id        string   `bson:"_id" validate:"required"`
	UserEmail string   `bson:"user_email" validate:"required,email"`
	Car       car      `bson:"car" validate:"required"`
	Position  position `bson:"position" validate:"required"`
}

type car struct {
	Title              string   `bson:"title" validate:"required"`
	Make               string   `bson:"make" validate:"required"`
	Model              string   `bson:"model" validate:"required"`
	Price              int      `bson:"price" validate:"required"`
	Mileage            int      `bson:"mileage" validate:"required"`
	Year               int      `bson:"year" validate:"required"`
	Photos             []string `bson:"photos" validate:"required"`
	Type               string   `bson:"type"`
	EngineCapacity     string   `bson:"engine_capacity"`
	Fuel               string   `bson:"fuel"`
	Power              int      `bson:"power"`
	Transmission       string   `bson:"transmission"`
	Drive              string   `bson:"drive"`
	Steering           string   `bson:"steering"`
	Doors              int      `bson:"doors"`
	Seats              int      `bson:"seats"`
	RegistrationNumber string   `bson:"registration_number"`
	FirstRegistration  int      `bson:"first_registration"`
	Condition          string   `bson:"condition"`
	VinNumber          int      `bson:"vin_number"`
	Description        string   `bson:"description"`
}

type position struct {
	PositionX float32 `bson:"position_x"`
	PositionY float32 `bson:"position_y"`
}
