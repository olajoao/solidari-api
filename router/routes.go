package router

import (
	"github.com/gin-gonic/gin"
	"github.com/olajoao/solidari-api/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/donations", handler.ListDonationsHandler)
		v1.GET("/donation/:id", handler.ShowDonationHandler)
		v1.POST("/donation", handler.CreateDonationHandler)
		v1.PUT("/donation/:id", handler.EditDonationHandler)
		v1.DELETE("/donation/:id", handler.DeleteDonationHandler)
	}
}
