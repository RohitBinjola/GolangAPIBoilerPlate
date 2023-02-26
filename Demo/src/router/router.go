package router

import (
	"Demo/src/api/controller"
	"Demo/src/config"
	"Demo/src/middleware"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := NewRouter()
	router.Run(config.Appconfig.GetString("server.port"))
}

func NewRouter() *gin.Engine {
	router := gin.New()
	resource := router.Group("/api")
	resource.Use(middleware.LogRequestInfo())
	{
		resource.GET("/GetData", controller.GetData)
		resource.GET("/GetQueryStringData", controller.GetQueryStringData)
	}
	return router
}
