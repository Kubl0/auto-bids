package controllers

import (
	"cars/models"
	"cars/responses"
	"cars/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func EditOffer(c *gin.Context) {
	result := make(chan responses.UserResponse)

	go func(cCp *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(result)
		var resultModel models.EditOffer
		validate := validator.New(validator.WithRequiredStructEnabled())
		email := models.Email{Email: cCp.Param("email")}

		if err := validate.Struct(email); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error validation email",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}
		if err := cCp.ShouldBindJSON(&resultModel); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error model edit_offer",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}
		if err := validate.Struct(resultModel); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error validation edit_offer",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}
		var userCollection = service.GetCollection(service.DB)
		objectId, err := primitive.ObjectIDFromHex(resultModel.Id)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Invalid Id",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}
		filter := bson.D{{"_id", objectId}, {"user_email", email.Email}}
		println(resultModel.Id)
		update := bson.D{
			{"$set", bson.D{
				{"car.description", resultModel.Description},
				{"car.price", resultModel.Price},
				{"car.mileage", resultModel.Mileage},
				{"car.photos", resultModel.Photos},
			}},
		}
		results, err := userCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error finding offer",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		if results.MatchedCount == 0 {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Document not found",
				Data:    map[string]interface{}{"error": results},
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
