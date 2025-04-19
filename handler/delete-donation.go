package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteDonationHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete donation",
	})
}
