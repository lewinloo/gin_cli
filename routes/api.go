package routes

import (
	"gin_cli/controller"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

func (r *RoleRouter) InitApiRoutes(Router *gin.RouterGroup) {
	router := Router.Group("api")
	controller := controller.Api{}
	{
		router.POST("create", controller.Create)    // 创建api
		router.POST("list", controller.QueryByPage) // 分页查询api
		router.GET("all", controller.QueryAll)      // 查询所有api
	}
}
