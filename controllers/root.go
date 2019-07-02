package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RootController - Placeholder struct for this controller
type RootController struct{}

// Index - Serve the / page of the site
func (r RootController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", gin.H{
		"heroTitle":    "Self Destruct",
		"heroSubtitle": "Send messages and secrets that self destruct after being viewed.",
	})
}

// About - About page
func (r RootController) About(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "about", gin.H{
		"heroTitle":    "About",
		"heroSubtitle": "",
	})
}

// Security - Security page
func (r RootController) Security(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "security", gin.H{
		"heroTitle": "Security",
	})
}
