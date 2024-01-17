package models

type Offer struct {
	Email    string `bson:"email" validate:"required,email"`
	Car      car
	Position position
}

type car struct {
	Title              string   `bson:"title" validate:"required"`
	Make               string   `bson:"make" validate:"required"`
	Model              string   `bson:"model" validate:"required"`
	Price              int      `bson:"price" validate:"required"`
	Mileage            int      `bson:"mileage" validate:"required"`
	Year               int      `bson:"year" validate:"required"`
	Type               string   `bson:"type" validate:"required"`
	Photos             []string `bson:"photos" validate:"required"`
	EngineCapacity     string   `bson:"engine_capacity" validate:"required"`
	Fuel               string   `bson:"fuel" validate:"required"`
	Power              int      `bson:"power" validate:"required"`
	Transmission       string   `bson:"transmission" validate:"required"`
	Drive              string   `bson:"drive" validate:"required"`
	Steering           string   `bson:"steering" validate:"required"`
	Doors              int      `bson:"doors" validate:"required"`
	Seats              int      `bson:"seats" validate:"required"`
	RegistrationNumber string   `bson:"registration_number" validate:"required"`
	FirstRegistration  int      `bson:"first_registration" validate:"required"`
	Condition          string   `bson:"condition" validate:"required"`
	VinNumber          int      `bson:"vin_number" validate:"required"`
	Description        string   `bson:"description" validate:"required"`
}

type position struct {
	PositionX float32 `bson:"position_x" validate:"required"`
	PositionY float32 `bson:"position_y" validate:"required"`
}
