package models

type CheckOffer struct {
	Make              string   `json:"make"`
	Model             string   `json:"model"`
	PriceMin          int      `json:"price_min"`
	PriceMax          int      `json:"price_max"`
	MileageMin        int      `json:"mileage_min"`
	MileageMax        int      `json:"mileage_max"`
	YearMin           int      `json:"year_min"`
	YearMax           int      `json:"year_max"`
	Type              string   `json:"type"`
	EngineCapacityMin int      `json:"engine_capacity_min"`
	EngineCapacityMax int      `json:"engine_capacity_max"`
	Fuel              string   `json:"fuel"`
	PowerMin          int      `json:"power_min"`
	PowerMax          int      `json:"power_max"`
	Transmission      string   `json:"transmission"`
	Drive             string   `json:"drive"`
	Steering          string   `json:"steering"`
	Doors             int      `json:"doors"`
	Seats             int      `json:"seats"`
	Condition         string   `json:"condition"`
	Location          Location `json:"location"`
	Distance          int      `json:"distance"`
}
