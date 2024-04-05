package controllers

import (
	"context"
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
				Status:  http.StatusBadRequest,
				Message: "Error validation profile",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		var userCollection = service.GetCollection(service.DB)
		existingProfile := models.PostProfile{}
		err := userCollection.FindOne(ctx, bson.M{"email": resultModel.Email}).Decode(&existingProfile)
		if err == nil {
			result <- responses.UserResponse{
				Status:  http.StatusOK,
				Message: "User logged",
				Data:    map[string]interface{}{"data": existingProfile},
			}
			return
		}
		results, err := userCollection.InsertOne(ctx, resultModel)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error adding profile",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}
		result <- responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "User created",
			Data:    map[string]interface{}{"data": results},
		}
	}(c.Copy())
	res := <-result
	c.JSON(res.Status, res)
}
