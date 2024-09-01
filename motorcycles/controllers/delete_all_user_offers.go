package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"motorcycles/models"
	"motorcycles/responses"
	"motorcycles/service"
	"net/http"
	"time"
)

func DeleteAllUserOffer(c *gin.Context) {
	result := make(chan responses.UserResponse)

	go func(cCp *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(result)
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

		var userCollection = service.GetCollection(service.DB)

		filter := bson.D{{"user_email", email.Email}}
		results, err := userCollection.DeleteMany(ctx, filter)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error finding offers",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}
		if results.DeletedCount < 1 {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Offer does not exist",
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
