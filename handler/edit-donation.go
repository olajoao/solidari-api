package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditDonationHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Edit donation",
	})
}
