package main

import (
	"github.com/gin-gonic/gin"
	"motorcycles/routes"
	"motorcycles/service"
	"os"
)

func main() {
	router := gin.Default()

	routes.MotorcyclesRoute(router)
	service.ConnectDB()

	errRouter := router.Run(os.Getenv("MOTORCYCLES_URI"))
	if errRouter != nil {
		panic(errRouter)
	}
}
