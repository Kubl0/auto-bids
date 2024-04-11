package test

import (
	"bytes"
	"cars/controllers"
	"cars/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostOffer(t *testing.T) {
	router := gin.Default()
	router.POST("/cars/add/:email", controllers.PostOffer)

	email := "test@example.com"
	car := models.Car{
		Title:       "Test Car",
		Make:        "Test Make",
		Model:       "Test Model",
		Price:       10000,
		Description: "Test Description",
		Photos:      []string{"https://example.com/image.jpg"},
		Year:        2020,
	}
	payload, _ := json.Marshal(car)

	req, err := http.NewRequest("POST", "/cars/add/"+email, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status %d; got %d", http.StatusCreated, resp.Code)
	}
}
