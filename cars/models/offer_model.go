package models

type Offer struct {
	Id        string `json:"_id" bson:"_id" validate:"required"`
	UserEmail string `json:"user_email" bson:"user_email" validate:"required,email"`
	Car       Car    `json:"car" bson:"car" validate:"required"`
}

type PostOffer struct {
	UserEmail string `json:"user_email" bson:"user_email" validate:"required,email"`
	Car       Car    `json:"car" bson:"car" validate:"required"`
}

type Car struct {
	Title              string   `json:"title" bson:"title" validate:"required"`
	Make               string   `json:"make" bson:"make" validate:"required"`
	Model              string   `json:"model" bson:"model" validate:"required"`
	Price              int      `json:"price" bson:"price" validate:"required"`
	Description        string   `json:"description" bson:"description" validate:"required"`
	Photos             []string `json:"photos" bson:"photos" validate:"required"`
	Year               int      `json:"year" bson:"year" validate:"required"`
	Mileage            int      `json:"mileage" bson:"mileage"`
	VinNumber          int      `json:"vin_number" bson:"vin_number"`
	EngineCapacity     string   `json:"engine_capacity" bson:"engine_capacity"`
	Fuel               string   `json:"fuel" bson:"fuel"`
	Transmission       string   `json:"transmission" bson:"transmission"`
	Steering           string   `json:"steering" bson:"steering"`
	Type               string   `json:"type" bson:"type"`
	Power              int      `json:"power" bson:"power"`
	Drive              string   `json:"drive" bson:"drive"`
	Doors              int      `json:"doors" bson:"doors"`
	Seats              int      `json:"seats" bson:"seats"`
	RegistrationNumber string   `json:"registration_number" bson:"registration_number"`
	FirstRegistration  string   `json:"first_registration" bson:"first_registration"`
	Condition          string   `json:"condition" bson:"condition"`
	TelephoneNumber    string   `json:"telephone_number" bson:"telephone_number"`
	Position           Position `json:"position" bson:"position"`
}

type Position struct {
	PositionX float32 `json:"position_x" bson:"position_x"`
	PositionY float32 `json:"position_y" bson:"position_y"`
}
