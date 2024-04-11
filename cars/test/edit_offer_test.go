package test

import (
	"bytes"
	"cars/controllers"
	"cars/models"
	"cars/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestEditOffer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var userCollection = service.GetCollection(service.DB)
	email := "test@test.pl"

	car := models.Car{
		Title:       "Test Car",
		Make:        "Test Make",
		Model:       "Test Model",
		Price:       10000,
		Description: "Test Description",
		Photos:      []string{"https://example.com/image.jpg"},
		Year:        2020,
	}

	newOffer := models.PostOffer{
		UserEmail: email,
		Car:       car,
	}
	results, err := userCollection.InsertOne(ctx, newOffer)
	if err != nil {
		t.Fatal(err)
	}

	router := gin.Default()
	router.PUT("/cars/edit/:email", controllers.EditOffer)

	newCar := models.EditOffer{
		Id:          results.InsertedID.(primitive.ObjectID).Hex(),
		Description: "new test description",
	}

	payload, _ := json.Marshal(newCar)

	req, err := http.NewRequest("PUT", "/cars/edit/"+email, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		fmt.Println(resp)
		t.Errorf("Expected status %d; got %d", http.StatusCreated, resp.Code)
	}
}
