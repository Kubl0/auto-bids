package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"
	"profiles/routes"

	_ "profiles/docs"
)

// @title Profiles API
// @version 1.0
// @description This is a simple CRUD API for profiles
// @host localhost:4100
// @BasePath /
func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SessionRoute(router)

	errRouter := router.Run(os.Getenv("PROFILES_URI"))
	if errRouter != nil {
		panic(errRouter)
	}
}
