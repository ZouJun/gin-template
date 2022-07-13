package group

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-template/controller"
)

func InitDemoRoute(demoGroup *gin.RouterGroup) {
	controller.DemoRoute(demoGroup)
	testGroup := demoGroup.Group("/test")
	{
		fmt.Println(testGroup)
	}
}
