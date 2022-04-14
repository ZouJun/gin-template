package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-template/controller"
	"go-gin-template/middleware"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.Cors())
	//router.LoadHTMLGlob("templates/*")

	//API
	apiGroup := router.Group("/api")
	{
		//接口
		demoGroup := apiGroup.Group("/demo")
		{
			controller.DemoRoute(demoGroup)
		}

	}

	return router
}
