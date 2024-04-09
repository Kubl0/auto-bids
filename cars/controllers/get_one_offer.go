package controllers

import (
	"cars/models"
	"cars/responses"
	"cars/service"
	"context"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetOneOffer godoc
// @Summary Get an offer
// @Description Get a user offer by id
// @ID get-offer
// @Produce json
// @Param id path string true "Id of the offer to be retrieved"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} responses.UserResponse
// @Failure 404 {object} responses.UserResponse
// @Failure 500 {object} responses.UserResponse
// @Router /cars/details/{id} [get]
func GetOneOffer(c *gin.Context) {
	result := make(chan responses.UserResponse)

	go func(cCp *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(result)
		var resultModel models.Offer
		validate := validator.New(validator.WithRequiredStructEnabled())

		offerId := models.Id{Id: cCp.Param("id")}
		if err := validate.Struct(offerId); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error validation id",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		objectId, err := primitive.ObjectIDFromHex(offerId.Id)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Invalid Id",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		var userCollection = service.GetCollection(service.DB)

		filter := bson.D{{"_id", objectId}}
		err = userCollection.FindOne(ctx, filter).Decode(&resultModel)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error finding offer",
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
