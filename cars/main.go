package main

import (
	"cars/routes"
	"cars/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	router := gin.Default()
	err := godotenv.Load()
	if err != nil {
		return
	}
	//routes.SessionRoute(router)
	routes.CarsRoute(router)
	service.ConnectDB()
	errRouter := router.Run(os.Getenv("PROFILES_URI"))
	if errRouter != nil {
		panic(errRouter)
	}
}
