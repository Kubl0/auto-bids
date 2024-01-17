package main

import (
	"cars/service"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()

	//routes.SessionRoute(router)
	service.ConnectDB()
	errRouter := router.Run(os.Getenv("PROFILES_URI"))
	if errRouter != nil {
		panic(errRouter)
	}
}
