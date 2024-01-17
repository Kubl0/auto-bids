package models

type GetOffer struct {
	Make           string `bson:"make"`
	Model          string `bson:"model"`
	Price          int    `bson:"price"`
	Mileage        int    `bson:"mileage"`
	Year           int    `bson:"year"`
	Type           string `bson:"type"`
	EngineCapacity string `bson:"engine_capacity"`
	Fuel           string `bson:"fuel"`
	Power          int    `bson:"power"`
	Transmission   string `bson:"transmission"`
	Drive          string `bson:"drive"`
	Steering       string `bson:"steering"`
	Doors          int    `bson:"doors"`
	Seats          int    `bson:"seats"`
	Condition      string `bson:"condition"`
}
