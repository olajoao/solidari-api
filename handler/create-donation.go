package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olajoao/solidari-api/schemas"
)

func CreateDonationHandler(ctx *gin.Context) {
	request := CreateDonationRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	donation := schemas.Donation{
		Title:       request.Title,
		Description: request.Description,
		Avatar:      request.Avatar,
		Location:    request.Location,
	}

	if err := db.Create(&donation).Error; err != nil {
		logger.Errorf("Error creating donation: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating donation on database")
		return
	}

	sendSuccess(ctx, "create-donation", donation)
}
