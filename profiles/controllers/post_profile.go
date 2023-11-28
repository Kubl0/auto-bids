package controllers

import (
	"context"
	"fmt"
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

		if err := cCp.ShouldBindJSON(&resultModel); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid request body",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		fmt.Print(resultModel)

		var userCollection = service.GetCollection(service.DB, "profiles")
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
