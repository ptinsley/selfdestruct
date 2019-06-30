package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RootController - Placeholder struct for this controller
type RootController struct{}

// Index - Serve the / page of the site
func (r RootController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", gin.H{})
}
