package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/donations", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Donations",
			})
		})

		v1.GET("/donation/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Donation",
			})
		})

		v1.POST("/donation", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "CREATE single donation",
			})
		})

		v1.PUT("/donation/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "EDIT single donation",
			})
		})

		v1.DELETE("/donation/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "DELETE Donation",
			})
		})
	}
}
