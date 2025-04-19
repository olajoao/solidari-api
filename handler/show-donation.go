package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowDonationHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Show single donation",
	})
}
