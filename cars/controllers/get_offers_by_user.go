package controllers

import (
	"cars/models"
	"cars/responses"
	"cars/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"time"
)

// GetOffersByUser godoc
// @Summary Get offers by user
// @Description Get a user offers by email
// @ID get-offers-by-user
// @Produce json
// @Param email path string true "Email address of the offers to be retrieved"
// @Param page path string true "Page number"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} responses.UserResponse
// @Failure 404 {object} responses.UserResponse
// @Failure 500 {object} responses.UserResponse
// @Router /cars/search/user/{email}/{page} [get]
func GetOffersByUser(c *gin.Context) {
	result := make(chan responses.UserResponse)

	go func(cCp *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(result)
		validate := validator.New(validator.WithRequiredStructEnabled())

		pageStr := cCp.Param("page")

		page, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error parse page",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		email := models.Email{Email: cCp.Param("email")}

		if err := validate.Struct(email); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error validation email",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		limit := int64(10)
		var userCollection = service.GetCollection(service.DB)

		filter := bson.M{"user_email": email.Email}
		opts := options.Find().SetSkip((page - 1) * 10).SetLimit(limit)
		results, err := userCollection.Find(ctx, filter, opts)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error finding offers",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		var offers []models.Offer
		if err := results.All(ctx, &offers); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error decoding offers",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		numberOfOffers, err := userCollection.CountDocuments(ctx, filter)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error counting offers",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		result <- responses.UserResponse{
			Status:  http.StatusOK,
			Message: "ok",
			Data:    map[string]interface{}{"data": offers, "number_of_pages": (numberOfOffers + 10 - 1) / 10},
		}

	}(c.Copy())
	res := <-result
	c.JSON(res.Status, res)
}
