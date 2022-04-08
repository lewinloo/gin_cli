package routes

import (
	"gin_cli/controller"

	"github.com/gin-gonic/gin"
)

type PermissionRouter struct {
}

func (r *RoleRouter) InitPermissionRoutes(Router *gin.RouterGroup) {
	router := Router.Group("permission")
	controller := controller.Permission{}
	{
		router.POST("update", controller.UpdatePermission) // 更新角色权限
	}
}
