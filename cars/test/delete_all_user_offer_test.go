package test

import (
	"cars/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteAllUserOffer(t *testing.T) {
	router := gin.Default()
	router.DELETE("/cars/delete/all/:email", controllers.DeleteAllUserOffer)

	req, err := http.NewRequest("DELETE", "/cars/delete/all/test@example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned unexpected error")
	}
}
