package controller

import (
	"gin_cli/service"
	"gin_cli/service/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

// @Tags 用户模块
// @Summary 用户注册接口
// @accept application/json
// @Produce application/json
// @Param params body request.RegisterParams true "注册参数"
// @Success 200 {object} response.Base "测试api成功返回体"
// @Router /user/register [post]
func (u *User) Register(c *gin.Context) {
	var params request.RegisterParams

	if err := c.ShouldBindJSON(&params); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    100001,
			"message": "参数格式错误",
			"success": false,
		})
	} else {
		if user, err := service.ServiceApp.UserService.Register(params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    100002,
				"message": err.Error(),
				"success": false,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "注册成功",
				"success": true,
				"data":    user,
			})
		}
	}
}

// @Tags 用户模块
// @Summary 用户登录接口
// @accept application/json
// @Produce application/json
// @Param params body request.RegisterParams true "登录参数"
// @Success 200 {object} response.Base "测试api成功返回体"
// @Router /user/login [post]
func (u *User) Login(c *gin.Context) {
	var params request.RegisterParams

	if err := c.ShouldBindJSON(&params); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    100001,
			"message": "参数格式错误",
			"success": false,
		})
	} else {
		if data, err := service.ServiceApp.UserService.Login(params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    100002,
				"message": err.Error(),
				"success": false,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "登录成功",
				"success": true,
				"data":    data,
			})
		}
	}
}
