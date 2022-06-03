package main

import (
	"github.com/gin-gonic/gin"
	"micro-wiki/middleware"
)

// 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	// 处理请求, 规范返回的中间件

	// 声明v1路由组
	v1Group := r.Group("/api/v1")
	// 定义v1路由组的所有业务接口
	v1User := v1Group.Group("/user")
	v1User.Use(middleware.JwtToken())
	{
		/* 用户模块 */
		// 注册用户
		//v1User.POST("/register", user.Register())
		// 登录
		//v1User.POST("/login", user.CheckLogin())
		// 修改用户信息
		//v1User.POST("/modify", nil)
	}

	return r
}
