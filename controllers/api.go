package controllers

import (
	"net/http"

	"github.com/ptinsley/selfdestruct/secret"
	"github.com/ptinsley/selfdestruct/utils"

	"github.com/gin-gonic/gin"
)

// APIController - placeholder struct for api controller
type APIController struct{}

// CreateRequest - capture the json fields we need from the post
type CreateRequest struct {
	Secret string `json:"secret"`
}

// Create - handle an inbound secret creation request
func (s APIController) Create(ctx *gin.Context) {
	var request CreateRequest
	ctx.BindJSON(&request)

	if request.Secret == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "'secret' parameter must be provided and not empty",
		})
		return
	}

	secretID, err := secret.Create(request.Secret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "currently unable to store secret, try again later",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"retreival_url": utils.RetrieveURL(ctx, secretID),
	})
}
