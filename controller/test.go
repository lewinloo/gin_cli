package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct {
}

// @Tags Test
// @Summary 测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param size query int false "当前页的大小"
// @Param params body interface{} false "xxx"
// @Success 200 {object} response.Base "测试api成功返回体"
// @Router /test/ping [post]
func (test *Test) GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
