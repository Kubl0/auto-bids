package test

import (
	"github.com/gin-gonic/gin"
	"motorcycles/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOffers(t *testing.T) {
	router := gin.Default()
	router.GET("/motorcycles/search/:page", controllers.GetOffers)

	req, err := http.NewRequest("GET", "/motorcycles/search/1?make=audi", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status":200,"message":"ok","data":{"data":null,"number_of_pages":0}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
