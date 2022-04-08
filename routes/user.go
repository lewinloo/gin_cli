package routes

import (
	"gin_cli/controller"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserRoutes(Router *gin.RouterGroup) {
	router := Router.Group("user")
	controller := controller.User{}
	{
		router.POST("register", controller.Register) // 注册
		router.POST("login", controller.Login)       // 注册
	}
}
