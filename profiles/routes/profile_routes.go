package routes

import (
	"github.com/gin-gonic/gin"
	"profiles/controllers"
)

func SessionRoute(router *gin.Engine) {
	profiles := router.Group("/profiles")
	{
		profiles.GET("/user_profile/:email", controllers.GetProfile)
		profiles.POST("/user_profile", controllers.PostProfile)
		profiles.DELETE("/user_profile/:email", controllers.DeleteProfile)
		profiles.PUT("/user_profile", controllers.EditProfile)
	}
}
