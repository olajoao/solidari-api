package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListDonationsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List donations from handler",
	})
}
