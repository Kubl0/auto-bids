package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"profiles/models"
	"profiles/responses"
	"profiles/service"
	"time"
)

func getProfile(c *gin.Context) {

	cCp := c.Copy()

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		var result models.Profile

		email := cCp.Param("email")
		var userCollection = service.GetCollection(service.DB, "profiles")
		err := userCollection.FindOne(ctx, bson.D{{"email", email}}).Decode(&result)
		if err != nil {
			cCp.JSON(
				http.StatusInternalServerError,
				responses.UserResponse{
					Status:  http.StatusInternalServerError,
					Message: "Error finding profile",
					Data:    map[string]interface{}{"error": err.Error()},
				})
		}

		cCp.JSON(
			http.StatusOK,
			responses.UserResponse{
				Status:  http.StatusOK,
				Message: "ok",
				Data:    map[string]interface{}{"data": result},
			},
		)
	}()
}
