package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController - struct for this controller
type HealthController struct{}

// Status - Return a 200 showing that our code is functional
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
