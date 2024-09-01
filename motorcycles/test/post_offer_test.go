package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"motorcycles/controllers"
	"motorcycles/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostOffer(t *testing.T) {
	router := gin.Default()
	router.POST("/motorcycles/add/:email", controllers.PostOffer)

	email := "test@example.com"
	motorcycle := models.Motorcycle{
		Title:       "Test Motorcycle",
		Make:        "Test Make",
		Model:       "Test Model",
		Price:       10000,
		Description: "Test Description",
		Photos:      []string{"https://example.com/image.jpg"},
		Year:        2020,
	}
	payload, _ := json.Marshal(motorcycle)

	req, err := http.NewRequest("POST", "/motorcycles/add/"+email, bytes.NewBuffer(payload))
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
