package controllers

import (
	"cars/responses"
	"cars/service"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

// DeleteOfferById godoc
// @Summary Delete all offers
// @Description Delete all offers by user email
// @ID delete-offer-by-id
// @Produce json
// @Param id path string true "ID of the offer to delete"
// @Success 200
// @Failure 400 {object} responses.UserResponse
// @Failure 404 {object} responses.UserResponse
// @Failure 500 {object} responses.UserResponse
// @Router /admin/cars/delete/{id} [delete]
func DeleteOfferById(c *gin.Context) {
	result := make(chan responses.UserResponse)

	go func(cCp *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(result)
		id := cCp.Param("id")

		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Invalid Id",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		var userCollection = service.GetCollection(service.DB, "cars")

		filter := bson.D{{"_id", objectId}}
		results, err := userCollection.DeleteOne(ctx, filter)
		if err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error finding offer",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}
		if results.DeletedCount != 1 {
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
