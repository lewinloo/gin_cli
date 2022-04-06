package utils

import (
	"github.com/gin-gonic/gin"
)

func GetRequestBody(c *gin.Context) string {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	return string(buf[0:n])
}
