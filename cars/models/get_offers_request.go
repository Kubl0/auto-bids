package models

type CheckOffer struct {
	Make              string   `json:"make" form:"make" validate:"max=30"`
	Model             string   `json:"model" form:"model" validate:"max=30"`
	PriceMin          int      `json:"price_min" form:"price_min"`
	PriceMax          int      `json:"price_max" form:"price_max"`
	MileageMin        int      `json:"mileage_min" form:"mileage_min"`
	MileageMax        int      `json:"mileage_max" form:"mileage_max"`
	YearMin           int      `json:"year_min" form:"year_min"`
	YearMax           int      `json:"year_max" form:"year_max"`
	Type              string   `json:"type" form:"type"`
	EngineCapacityMin int      `json:"engine_capacity_min" form:"engine_capacity_min"`
	EngineCapacityMax int      `json:"engine_capacity_max" form:"engine_capacity_max"`
	Fuel              string   `json:"fuel" form:"fuel"`
	PowerMin          int      `json:"power_min" form:"power_min"`
	PowerMax          int      `json:"power_max" form:"power_max"`
	Transmission      string   `json:"transmission" form:"transmission"`
	Drive             string   `json:"drive" form:"drive"`
	Steering          string   `json:"steering" form:"steering"`
	Doors             int      `json:"doors" form:"doors"`
	Seats             int      `json:"seats" form:"seats"`
	Condition         string   `json:"condition" form:"condition"`
	Location          Location `json:"location" form:"location"`
	Distance          int      `json:"distance" form:"distance" validate:"max=1000"`
}
