package main

import (
	"cars/routes"
	"cars/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"

	_ "cars/docs"
)

// @title Profiles API
// @version 1.0
// @description This is a simple CRUD API for profiles
// @host localhost:4200
// @BasePath /
func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.CarsRoute(router)
	service.ConnectDB()
	errRouter := router.Run(os.Getenv("CARS_URI"))
	if errRouter != nil {
		panic(errRouter)
	}
}
