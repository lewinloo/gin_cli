package routes

import (
	"gin_cli/controller"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct {
}

func (r *RoleRouter) InitRoleRoutes(Router *gin.RouterGroup) {
	router := Router.Group("role")
	controller := controller.Role{}
	{
		router.POST("create", controller.Create) // 创建角色
	}
}
