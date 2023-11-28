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

func EditProfile(c *gin.Context) {
	result := make(chan responses.UserResponse)

	go func(cCp *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(result)
		var resultModel models.PostProfile

		if err := cCp.BindJSON(&resultModel); err != nil {
			result <- responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error model edit_profile",
				Data:    map[string]interface{}{"error": err.Error()},
			}
			return
		}

		var userCollection = service.GetCollection(service.DB, "profiles")

		filter := bson.D{{"email", resultModel.Email}}
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
