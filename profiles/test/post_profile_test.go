package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"profiles/controllers"
	"profiles/models"
	"testing"
)

func TestPostProfile(t *testing.T) {
	router := gin.Default()
	router.GET("/profiles/user", controllers.PostProfile)

	profile := models.PostProfile{
		UserName:     "Test_user",
		Email:        "test@example.com",
		ProfileImage: "https://example.com/image.jpg",
	}
	payload, _ := json.Marshal(profile)

	req, err := http.NewRequest("GET", "/profiles/user", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusCreated, resp.Code)
	}
}
