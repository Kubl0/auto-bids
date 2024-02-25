package routes

import (
	"cars/controllers"
	"github.com/gin-gonic/gin"
)

func CarsRoute(router *gin.Engine) {
	profiles := router.Group("/cars")
	{
		profiles.GET("/offer/:id", controllers.GetOneOffer)
		profiles.GET("/user/:email/:page", controllers.GetOffersByUser)
		profiles.POST("/page/:page", controllers.GetOffers)
		profiles.POST("/add/:email", controllers.PostOffer)
		profiles.DELETE("/delete/:email", controllers.DeleteOffer)
		profiles.DELETE("/delete/all/:email", controllers.DeleteAllUserOffer)
		profiles.PUT("/edit/:email", controllers.EditOffer)
	}
}
