package main

import (
	"cars/routes"
	"cars/service"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()

	routes.CarsRoute(router)
	service.ConnectDB()
	errRouter := router.Run(os.Getenv("CARS_URI"))
	if errRouter != nil {
		panic(errRouter)
	}
}
