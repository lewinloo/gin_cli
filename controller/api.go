package controller

import (
	"gin_cli/service"
	"gin_cli/service/request"
	"gin_cli/utils/response"

	"github.com/gin-gonic/gin"
)

type Api struct{}

// @Tags Api管理模块
// @Summary 创建一个Api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param params body request.CreateApiParams true "创建api参数"
// @Success 200 {object} response.Base "成功返回体"
// @Router /api/create [post]
func (a *Api) Create(c *gin.Context) {
	var params request.CreateApiParams
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(400, "参数错误", c)
	} else if err := service.ServiceApp.ApiService.CreateApi(params); err != nil {
		response.Fail(500, err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Api管理模块
// @Summary 分页且可以添加条件的查询api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param params body request.QueryApiListParams true "分页查询api参数"
// @Success 200 {object} response.Base "成功返回体"
// @Router /api/list [post]
func (a *Api) QueryByPage(c *gin.Context) {
	var params request.QueryApiListParams
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(400, "参数错误", c)
	} else if list, total, cur, size, err := service.ServiceApp.ApiService.GetApiList(params); err != nil {
		response.Fail(500, err.Error(), c)
	} else {
		response.OkWithListData(list, total, cur, size, c)
	}
}

// @Tags Api管理模块
// @Summary 查询所有api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Base "成功返回体"
// @Router /api/all [get]
func (a *Api) QueryAll(c *gin.Context) {
	if apis, err := service.ServiceApp.ApiService.GetAllApis(); err != nil {
		response.Fail(500, err.Error(), c)
	} else {
		response.OkWithData("查询成功", apis, c)
	}
}
