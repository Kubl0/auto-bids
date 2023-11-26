package controllers

import (
	"context"
	"net/http"
	"profiles/models"
	"profiles/responses"
	"profiles/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProfile(c *gin.Context) {
	result := make(chan responses.UserResponse)

	go func(cCp *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(result)
		var resultModel models.Profile

		email := cCp.Param("email")
		var userCollection = service.GetCollection(service.DB, "profiles")
		err := userCollection.FindOne(ctx, bson.D{{"email", email}}).Decode(&resultModel)
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
