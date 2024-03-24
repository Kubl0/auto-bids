package controllers

import (
	"cars/models"
	"cars/responses"
	"cars/service"
	"context"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// PostOffer godoc
// @Summary Post an offer
// @Description Post a user offer
// @ID post-offer
// @Produce json
// @Param email path string true "Email address of the offer to be posted"
// @Param userData body models.Car true "User data to be posted"
// @Success 201 {object} responses.UserResponse
// @Failure 400 {object} responses.UserResponse
// @Failure 500 {object} responses.UserResponse
// @Router /cars/add/{email} [post]
func PostOffer(c *gin.Context) {
	result := make(chan responses.UserResponse)

	go func(cCp *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(result)
		var resultCar models.Car
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

		if err := cCp.ShouldBindJSON(&resultCar); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid request body",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		if err := validate.Struct(resultCar); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Error validation car",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		var userCollection = service.GetCollection(service.DB, "cars")
		newOffer := models.PostOffer{
			UserEmail: email.Email,
			Car:       resultCar,
		}
		results, err := userCollection.InsertOne(ctx, newOffer)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error adding offer",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}
		result <- responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "Offer created",
			Data:    map[string]interface{}{"data": results},
		}
	}(c.Copy())
	res := <-result
	c.JSON(res.Status, res)
}
