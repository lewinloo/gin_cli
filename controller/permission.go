package controller

import (
	"gin_cli/service"
	"gin_cli/service/request"
	"gin_cli/utils/response"

	"github.com/gin-gonic/gin"
)

type Permission struct{}

// @Tags 权限管理
// @Summary 更新角色权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PermissionInReceive true "权限id, 权限模型列表"
// @Success 200 {object} response.Base "更新角色api权限"
// @Router /permission/update [post]
func (p *Permission) UpdatePermission(c *gin.Context) {
	var params request.PermissionInReceive
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(400, "参数错误", c)
	} else if err := service.ServiceApp.CasbinService.UpdateCasbin(params.RoleId, params.PermissionInfos); err != nil {
		response.Fail(500, err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
