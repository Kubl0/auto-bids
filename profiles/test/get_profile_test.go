package test

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"profiles/controllers"
	"profiles/models"
	"profiles/service"
	"testing"
	"time"
)

func TestGetProfile(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var userCollection = service.GetCollection(service.DB)
	email := "test@test.pl"

	newProfile := models.PostProfile{
		UserName:     "test_user",
		Email:        email,
		ProfileImage: "https://example.com/image.jpg",
	}

	_, err := userCollection.InsertOne(ctx, newProfile)
	if err != nil {
		t.Fatal(err)
	}

	router := gin.Default()
	router.GET("/profiles/user/:email", controllers.GetProfile)

	req, err := http.NewRequest("GET", "/profiles/user/"+email, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
