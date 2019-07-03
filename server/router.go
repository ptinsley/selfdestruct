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

	// html/template with master templates
	router.HTMLRender = ginview.Default()

	// serve static files from the assets dir
	router.Static("/assets", "./assets")

	rootController := new(controllers.RootController)
	router.GET("/", rootController.Index)
	router.GET("/about", rootController.About)
	router.GET("/security", rootController.Security)

	healthController := new(controllers.HealthController)
	router.GET("/healthz", healthController.Status)

	secretController := new(controllers.SecretController)
	secretGroup := router.Group("secret")
	{
		secretGroup.GET("/", secretController.New)
		secretGroup.POST("/create", secretController.Create)
		secretGroup.GET("/retrieve/:id", secretController.Retrieve)
		secretGroup.GET("/reveal/:id", secretController.Reveal)

	}

	apiController := new(controllers.APIController)
	apiGroup := router.Group("api/v1")
	{
		apiGroup.POST("/create", apiController.Create)
	}

	return router
}
