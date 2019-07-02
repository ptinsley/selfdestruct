package controllers

import (
	"net/http"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/ptinsley/selfdestruct/secret"
	"github.com/ptinsley/selfdestruct/utils"
)

// SecretController - placeholder struct for secret controller
type SecretController struct{}

// New - display the secret creation page
func (s SecretController) New(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "secret/index", gin.H{
		"heroTitle":    "Create Secret",
		"heroSubtitle": "Create a one time secret",
	})
}

// Create - store the submitted secret
func (s SecretController) Create(ctx *gin.Context) {
	secretValue := ctx.PostForm("secret")
	url := location.Get(ctx)

	uuid := utils.UUID()

	if err := secret.Create(uuid, secretValue); err != nil {
		ctx.HTML(http.StatusInternalServerError, "secret/create", gin.H{
			"flashFailure": "Cannot currently store your secret, please try again later.",
		})
	} else {
		ctx.HTML(http.StatusOK, "secret/create", gin.H{
			"flashSuccess": "Secret Created!",
			"heroTitle":    "Share Secret",
			"heroSubtitle": "Now that you've created your secret, it's time to send it to someone.",
			"scheme":       url.Scheme,
			"host":         ctx.Request.Host,
			"uuid":         uuid,
		})
	}

}

// Retrieve - get the secret, delete the encryption key and serve the secret to the user
func (s SecretController) Retrieve(ctx *gin.Context) {
	id := ctx.Param("id")
	secret, _ := secret.Take(id)

	// FIXME: add error page/flash if the secret is gone

	ctx.HTML(http.StatusOK, "secret/retrieve", gin.H{
		"heroTitle":    "Your Secret",
		"heroSubtitle": "This message will self destruct when you leave the page.",
		"hideNav":      1,
		"secret":       secret,
	})
}
