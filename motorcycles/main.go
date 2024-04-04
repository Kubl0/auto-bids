package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()

	errRouter := router.Run(os.Getenv("MOTORCYCLES_URI"))
	if errRouter != nil {
		panic(errRouter)
	}
}
