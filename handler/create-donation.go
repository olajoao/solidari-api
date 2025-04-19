package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDonationHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "CREATE single donation",
	})
}
