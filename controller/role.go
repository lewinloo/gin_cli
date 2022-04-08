package controller

import (
	"gin_cli/service"
	"gin_cli/service/request"
	"gin_cli/utils/response"

	"github.com/gin-gonic/gin"
)

type Role struct{}

// @Tags 角色管理模块
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param params body request.CreateRoleParams true "创建角色参数"
// @Success 200 {object} response.Base "成功返回体"
// @Router /role/create [post]
func (r *Role) Create(c *gin.Context) {
	var params request.CreateRoleParams
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(400, "参数错误", c)
	} else if _, err := service.ServiceApp.RoleService.CreateRole(params); err != nil {
		response.Fail(500, err.Error(), c)
	} else {
		// 给角色添加默认权限
		service.ServiceApp.CasbinService.UpdateCasbin(params.RoleId, request.DefaultCasbin())
		response.OkWithMessage("创建成功", c)
	}
}
