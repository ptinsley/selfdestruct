package server

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"

	"github.com/ptinsley/selfdestruct/controllers"
)

// NewRouter - Build the router for the application
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// surface information about the request host, path and scheme in the gin.ctx
	router.Use(location.Default())

	router.HTMLRender = ginview.Default()

	root := new(controllers.RootController)
	router.GET("/", root.Index)

	health := new(controllers.HealthController)
	router.GET("/healthz", health.Status)

	secretGroup := router.Group("secret")
	{
		secretController := new(controllers.SecretController)
		secretGroup.GET("/", secretController.New)
		secretGroup.POST("/create", secretController.Create)
		secretGroup.GET("retrieve/:id", secretController.Retrieve)
	}

	return router
}
