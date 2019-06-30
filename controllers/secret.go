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
	ctx.HTML(http.StatusOK, "secret/index", gin.H{})
}

// Create - store the submitted secret
func (s SecretController) Create(ctx *gin.Context) {
	secretValue := ctx.PostForm("secret")
	url := location.Get(ctx)

	uuid := utils.UUID()
	secret.Create(uuid, secretValue)

	ctx.HTML(http.StatusOK, "secret/create", gin.H{
		"scheme": url.Scheme,
		"host":   url.Host,
		"uuid":   uuid,
	})
}

// Retrieve - get the secret, delete the encryption key and serve the secret to the user
func (s SecretController) Retrieve(ctx *gin.Context) {
	id := ctx.Param("id")
	secret, _ := secret.Take(id)

	ctx.HTML(http.StatusOK, "secret/retrieve", gin.H{
		"secret": secret,
	})
}
