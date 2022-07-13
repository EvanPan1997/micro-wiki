package router

import (
	"github.com/gin-gonic/gin"
	"micro-wiki/router/v1/system"
)

func InitRoute(r *gin.Engine) {
	// 初始化路由
	v1 := r.Group("/v1")

	var userRouter system.UserRouter
	userRouter.InitUserRouter(v1)
}
