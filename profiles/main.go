package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"profiles/routes"
)

func main() {
	router := gin.Default()

	routes.SessionRoute(router)

	errRouter := router.Run(os.Getenv("PROFILES_URI"))
	if errRouter != nil {
		panic(errRouter)
	}
}
