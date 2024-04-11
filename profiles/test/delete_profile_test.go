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

func TestDeleteProfile(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var userCollection = service.GetCollection(service.DB)
	email := "del_test@test.pl"

	profile := models.PostProfile{
		UserName:     "del_test_user",
		Email:        email,
		ProfileImage: "https://example.com/image.jpg",
	}

	_, err := userCollection.InsertOne(ctx, profile)
	if err != nil {
		t.Fatal(err)
	}

	router := gin.Default()
	router.DELETE("/profiles/delete/:email", controllers.DeleteProfile)

	req, err := http.NewRequest("DELETE", "/profiles/delete/"+email, nil)
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
