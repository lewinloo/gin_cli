package middleware

import (
	"fmt"
	"gin_cli/service"
	"gin_cli/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var casbinService = service.ServiceApp.CasbinService

// 拦截器
func CheckPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		waitUse, _ := utils.ParseToken(token)
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.RoleId
		e := casbinService.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)

		fmt.Println("sub => ", waitUse)

		// 规定角色id为777是超级管理员
		if success || sub == "777" {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"message": "权限不足", "success": false})
			c.Abort()
			return
		}
	}
}
