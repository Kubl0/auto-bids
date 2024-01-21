package routes

import (
	"cars/controllers"
	"github.com/gin-gonic/gin"
)

func CarsRoute(router *gin.Engine) {
	profiles := router.Group("/cars")
	{
		profiles.POST("/page/:page", controllers.GetOffers)
		profiles.POST("/add/:email", controllers.PostOffer)
		//profiles.DELETE("/user/:email", controllers.DeleteProfile)
		//profiles.PUT("/user/:email", controllers.EditProfile)
	}
}
