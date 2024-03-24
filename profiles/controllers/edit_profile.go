package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"profiles/models"
	"profiles/responses"
	"profiles/service"
	"time"
)

// EditProfile godoc
// @Summary Edit a profile
// @Description Edit a user profile by email
// @ID edit-profile
// @Produce json
// @Param email path string true "Email address of the profile to be edited"
// @Param userData body models.EditProfile true "User data to be edited"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} responses.UserResponse
// @Failure 404 {object} responses.UserResponse
// @Failure 500 {object} responses.UserResponse
// @Router /profiles/user/{email} [put]
func EditProfile(c *gin.Context) {
	result := make(chan responses.UserResponse)

	go func(cCp *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(result)
		var resultModel models.EditProfile
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

		if err := cCp.ShouldBindJSON(&resultModel); err != nil {
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
				Message: "Error validation profile",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		var userCollection = service.GetCollection(service.DB, "profiles")

		filter := bson.D{{"email", email.Email}}
		update := bson.D{
			{"$set", bson.D{
				{"profile_image", resultModel.ProfileImage},
				{"user_name", resultModel.UserName},
			}},
		}

		results, err := userCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error finding profile",
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
