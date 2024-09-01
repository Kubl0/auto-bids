package queries

import (
	"go.mongodb.org/mongo-driver/bson"
	"motorcycles/models"
)

func GetOfferQuery(offer models.CheckOffer) bson.M {
	query := bson.M{}

	if offer.Make != "" {
		query["motorcycle.make"] = offer.Make
	}

	if offer.Model != "" {
		query["motorcycle.model"] = offer.Model
	}

	if offer.PriceMin != 0 && offer.PriceMax != 0 {
		query["motorcycle.price"] = bson.M{"$gte": offer.PriceMin, "$lte": offer.PriceMax}
	} else if offer.PriceMin != 0 {
		query["motorcycle.price"] = bson.M{"$gte": offer.PriceMin}
	} else if offer.PriceMax != 0 {
		query["motorcycle.price"] = bson.M{"$lte": offer.PriceMax}
	}

	if offer.MileageMin != 0 && offer.MileageMax != 0 {
		query["motorcycle.mileage"] = bson.M{"$gte": offer.MileageMin, "$lte": offer.MileageMax}
	} else if offer.MileageMin != 0 {
		query["motorcycle.mileage"] = bson.M{"$gte": offer.MileageMin}
	} else if offer.MileageMax != 0 {
		query["motorcycle.mileage"] = bson.M{"$lte": offer.MileageMax}
	}

	if offer.YearMin != 0 && offer.YearMax != 0 {
		query["motorcycle.year"] = bson.M{"$gte": offer.YearMin, "$lte": offer.YearMax}
	} else if offer.YearMin != 0 {
		query["motorcycle.year"] = bson.M{"$gte": offer.YearMin}
	} else if offer.YearMax != 0 {
		query["motorcycle.year"] = bson.M{"$lte": offer.YearMax}
	}

	if offer.Type != "" {
		query["motorcycle.type"] = offer.Type
	}

	if offer.EngineCapacityMin != 0 && offer.EngineCapacityMax != 0 {
		query["motorcycle.engine_capacity"] = bson.M{"$gte": offer.EngineCapacityMin, "$lte": offer.EngineCapacityMax}
	} else if offer.EngineCapacityMin != 0 {
		query["motorcycle.engine_capacity"] = bson.M{"$gte": offer.EngineCapacityMin}
	} else if offer.EngineCapacityMax != 0 {
		query["motorcycle.engine_capacity"] = bson.M{"$lte": offer.EngineCapacityMax}
	}

	if offer.Fuel != "" {
		query["motorcycle.fuel"] = offer.Fuel
	}

	if offer.PowerMin != 0 && offer.PowerMax != 0 {
		query["motorcycle.power"] = bson.M{"$gte": offer.PowerMin, "$lte": offer.PowerMax}
	} else if offer.PowerMin != 0 {
		query["motorcycle.power"] = bson.M{"$gte": offer.PowerMin}
	} else if offer.PowerMax != 0 {
		query["motorcycle.power"] = bson.M{"$lte": offer.PowerMax}
	}

	if offer.Transmission != "" {
		query["motorcycle.transmission"] = offer.Transmission
	}

	if offer.DriveType != "" {
		query["motorcycle.drive_type"] = offer.DriveType
	}

	if offer.Condition != "" {
		query["motorcycle.condition"] = offer.Condition
	}

	if offer.CoordinatesX != 0 && offer.CoordinatesY != 0 && offer.Distance != 0 {
		query["motorcycle.location"] = bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{
					[]interface{}{offer.CoordinatesX, offer.CoordinatesY},
					offer.Distance / 6378100, // distance is in kilometers
				},
			},
		}
	}

	return query
}
