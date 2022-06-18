package middleware

import (
	"github.com/gin-gonic/gin"
	"micro-wiki/util/msg"
	"net/http"
)

// Jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		code := msg.SUCCESS
		if token == "" {
			// 若为空, 则203重定向
			code = msg.REDIRECT
			c.Abort()
		}
		/*
			解析token:
			1. 验证是否在数据库中可以找到
			2. 用盐值解析
			3. 是否超时
			4. 是否密码未被修改
		*/
		// 验证是否在数据库中可以找到
		// 用盐值解析
		// 是否超时
		// 是否密码未被修改

		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": msg.GetMsg(code),
		})
		//c.Set("username", key.Username)
		c.Next()
	}
}
