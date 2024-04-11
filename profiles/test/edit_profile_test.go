package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"profiles/controllers"
	"profiles/models"
	"profiles/service"
	"testing"
	"time"
)

func TestEditProfile(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var userCollection = service.GetCollection(service.DB)
	email := "test@test.pl"

	profile := models.PostProfile{
		UserName:     "test_user",
		Email:        email,
		ProfileImage: "https://example.com/image.jpg",
	}

	_, err := userCollection.InsertOne(ctx, profile)
	if err != nil {
		t.Fatal(err)
	}

	router := gin.Default()
	router.PUT("/profiles/user/:email", controllers.EditProfile)

	newProfile := models.EditProfile{
		UserName:     "new_test_user",
		ProfileImage: "https://example.com/image.jpg",
	}

	payload, _ := json.Marshal(newProfile)

	req, err := http.NewRequest("PUT", "/profiles/user/"+email, bytes.NewBuffer(payload))
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
