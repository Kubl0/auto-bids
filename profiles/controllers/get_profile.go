package controllers

import (
	"context"
	"github.com/go-playground/validator/v10"
	"net/http"
	"profiles/models"
	"profiles/responses"
	"profiles/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetProfile godoc
// @Summary Get a profile
// @Description Get a user profile by email
// @ID get-profile
// @Produce json
// @Param email path string true "Email address of the profile to be retrieved"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} responses.UserResponse
// @Failure 404 {object} responses.UserResponse
// @Failure 500 {object} responses.UserResponse
// @Router /profiles/user/{email} [get]
func GetProfile(c *gin.Context) {
	result := make(chan responses.UserResponse)

	go func(cCp *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(result)
		var resultModel models.Profile
		validate := validator.New(validator.WithRequiredStructEnabled())

		email := models.Email{Email: cCp.Param("email")}

		if err := validate.Struct(email); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error validation profile",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		var userCollection = service.GetCollection(service.DB, "profiles")
		err := userCollection.FindOne(ctx, bson.D{{"email", email.Email}}).Decode(&resultModel)
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
			Data:    map[string]interface{}{"data": resultModel},
		}
	}(c.Copy())
	res := <-result
	c.JSON(res.Status, res)
}
