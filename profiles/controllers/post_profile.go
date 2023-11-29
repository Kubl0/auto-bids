package controllers

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"profiles/models"
	"profiles/responses"
	"profiles/service"
	"time"

	"github.com/gin-gonic/gin"
)

func PostProfile(c *gin.Context) {
	result := make(chan responses.UserResponse)

	go func(cCp *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(result)
		var resultModel models.PostProfile
		validate := validator.New(validator.WithRequiredStructEnabled())

		fmt.Println(resultModel)

		if err := cCp.ShouldBindJSON(&resultModel); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid request body",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		if err := validate.Struct(resultModel); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error validation profile",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		var userCollection = service.GetCollection(service.DB, "profiles")
		existingProfile := models.PostProfile{}
		err := userCollection.FindOne(ctx, bson.M{"email": resultModel.Email}).Decode(&existingProfile)
		if err == nil {
			result <- responses.UserResponse{
				Status:  http.StatusConflict,
				Message: "Profile with the given email already exists",
				Data:    map[string]interface{}{"error": "Email already in use"},
			}
			return
		}
		results, err := userCollection.InsertOne(ctx, resultModel)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error finding profile",
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
