package routes

import (
	"gin_cli/controller"
	"github.com/gin-gonic/gin"
)

type TestRouter struct {
}

func (r *TestRouter) InitTestRoutes(Router *gin.RouterGroup) {
	testRouter := Router.Group("test")
	controller := controller.Test{}
	{
		testRouter.GET("ping", controller.GetHello) // 测试
	}
}
