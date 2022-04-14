package controller

import "github.com/gin-gonic/gin"

type DemoController struct {
}

//路由
func DemoRoute(router *gin.RouterGroup) {
	demoController := DemoController{}
	router.GET("/demo", demoController.demo)
}

func (*DemoController) demo(c *gin.Context) {

}
