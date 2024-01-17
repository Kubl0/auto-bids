package controllers

import (
	"cars/models"
	"cars/responses"
	"cars/service"
	"context"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func buildQuery(offer models.GetOffer) bson.M {
	query := bson.M{}

	if offer.Make != "" {
		query["make"] = offer.Make
	}
	if offer.Model != "" {
		query["model"] = offer.Model
	}
	if offer.Price != 0 {
		query["price"] = offer.Price
	}
	if offer.Mileage != 0 {
		query["mileage"] = offer.Mileage
	}
	if offer.Year != 0 {
		query["year"] = offer.Year
	}
	if offer.Type != "" {
		query["type"] = offer.Type
	}
	if offer.EngineCapacity != "" {
		query["engine_capacity"] = offer.EngineCapacity
	}
	if offer.Fuel != "" {
		query["fuel"] = offer.Fuel
	}
	if offer.Power != 0 {
		query["power"] = offer.Power
	}
	if offer.Transmission != "" {
		query["transmission"] = offer.Transmission
	}
	if offer.Drive != "" {
		query["drive"] = offer.Drive
	}
	if offer.Steering != "" {
		query["steering"] = offer.Steering
	}
	if offer.Doors != 0 {
		query["doors"] = offer.Doors
	}
	if offer.Seats != 0 {
		query["seats"] = offer.Seats
	}
	if offer.Condition != "" {
		query["condition"] = offer.Condition
	}

	return query
}

func GetOffers(c *gin.Context) {
	result := make(chan responses.UserResponse)

	go func(cCp *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(result)
		var resultModel models.GetOffer
		validate := validator.New(validator.WithRequiredStructEnabled())

		pageStr := cCp.Param("page")
		page, err := strconv.ParseInt(pageStr, 10, 64)

		if err := cCp.BindJSON(&resultModel); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error model edit_profile",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		if err := validate.Struct(resultModel); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error validation offer",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		limit := int64(10)

		var userCollection = service.GetCollection(service.DB, "profiles")
		filter := buildQuery(resultModel)
		opts := options.Find().SetSkip(page).SetLimit(limit)

		results, err := userCollection.Find(ctx, filter, opts)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error finding cars",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}
		result <- responses.UserResponse{
			Status:  http.StatusOK,
			Message: "ok",
			Data:    map[string]interface{}{"data": results},
		}

	}(c.Copy())
	res := <-result
	c.JSON(res.Status, res)
}
