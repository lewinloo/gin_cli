package routes

import (
	"gin_cli/controller"

	"github.com/gin-gonic/gin"
)

type PublicRouter struct {
}

func (r *UserRouter) InitPublicRoutes(Router *gin.RouterGroup) {
	router := Router.Group("")
	controller := controller.User{}
	{
		router.POST("/user/login", controller.Login) // 注册
	}
}
