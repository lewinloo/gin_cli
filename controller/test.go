package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Test struct {
}

// @Tags Test
// @Summary 测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Base "测试api成功返回体"
// @Router /test/ping [get]
func (test *Test) GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
