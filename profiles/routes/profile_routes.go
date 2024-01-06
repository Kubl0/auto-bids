package routes

import (
	"github.com/gin-gonic/gin"
	"profiles/controllers"
)

func SessionRoute(router *gin.Engine) {
	profiles := router.Group("/profiles")
	{
		profiles.GET("/user/:email", controllers.GetProfile)
		profiles.GET("/user", controllers.PostProfile)
		profiles.DELETE("/user/:email", controllers.DeleteProfile)
		profiles.PUT("/user", controllers.EditProfile)
	}
}
