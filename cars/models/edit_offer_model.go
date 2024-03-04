package models

type EditOffer struct {
	Id          string   `json:"id" bson:"_id" validate:"required"`
	Description string   `json:"description" bson:"description" validate:"max=500"`
	Price       int      `json:"price" bson:"price"`
	Mileage     int      `json:"mileage" bson:"mileage"`
	Photos      []string `json:"photos" bson:"photos"`
}
