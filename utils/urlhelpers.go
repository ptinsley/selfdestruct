package utils

import (
	"fmt"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

// RetrieveURL - build the url to retrieve a given secret
func RetrieveURL(ctx *gin.Context, secretID string) string {
	url := location.Get(ctx)

	return fmt.Sprintf("%s://%s/secret/retrieve/%s", url.Scheme, ctx.Request.Host, secretID)
}
