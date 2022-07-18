package router

import (
	"github.com/gin-gonic/gin"
	"micro-wiki/midderware/common"
	"micro-wiki/router/v1/system"
)

func InitRoute(r *gin.Engine) {
	// 初始化路由
	r.Use(common.Cors())
	v1 := r.Group("/v1")

	var userRouter system.UserRouter
	userRouter.InitUserRouter(v1)
}
