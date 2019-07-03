package utils

import (
	"fmt"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

// RetrieveURL - build the url for the retrieve page
func RetrieveURL(ctx *gin.Context, secretID string) string {
	url := location.Get(ctx)

	return fmt.Sprintf("%s://%s/secret/retrieve/%s", url.Scheme, ctx.Request.Host, secretID)
}

// RevealURL - build the url to reveal a given secret
func RevealURL(ctx *gin.Context, secretID string) string {
	url := location.Get(ctx)

	return fmt.Sprintf("%s://%s/secret/reveal/%s", url.Scheme, ctx.Request.Host, secretID)
}
