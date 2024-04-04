package models

type CheckOffer struct {
	Make              string  `json:"make" form:"make" validate:"max=30"`
	Model             string  `json:"model" form:"model" validate:"max=30"`
	PriceMin          int     `json:"price_min" form:"price_min"`
	PriceMax          int     `json:"price_max" form:"price_max"`
	MileageMin        int     `json:"mileage_min" form:"mileage_min"`
	MileageMax        int     `json:"mileage_max" form:"mileage_max"`
	YearMin           int     `json:"year_min" form:"year_min"`
	YearMax           int     `json:"year_max" form:"year_max"`
	Type              string  `json:"type" form:"type"`
	EngineCapacityMin int     `json:"engine_capacity_min" form:"engine_capacity_min"`
	EngineCapacityMax int     `json:"engine_capacity_max" form:"engine_capacity_max"`
	Fuel              string  `json:"fuel" form:"fuel"`
	PowerMin          int     `json:"power_min" form:"power_min"`
	PowerMax          int     `json:"power_max" form:"power_max"`
	Transmission      string  `json:"transmission" form:"transmission"`
	DriveType         string  `json:"drive_type" form:"drive_type"`
	Condition         string  `json:"condition" form:"condition"`
	CoordinatesX      float32 `json:"lat" form:"lat"`
	CoordinatesY      float32 `json:"lng" form:"lng"`
	Distance          float64 `json:"radius" form:"radius"`
	FilterBy          string  `json:"filter_by" form:"filter_by"`
	SortDirection     int     `json:"sort_direction" form:"sort_direction"`
}
