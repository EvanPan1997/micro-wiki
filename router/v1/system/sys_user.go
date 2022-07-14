package system

import (
	"github.com/gin-gonic/gin"
	v1 "micro-wiki/api/v1"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userGroup := router.Group("/user")

	{
		userGroup.POST("/userRegister", v1.ApiGroupApp.SystemApiGroup.Register) // 管理员注册用户
		userGroup.DELETE("/deleteUser")                                         // 管理员注销用户
		userGroup.POST("/setUserAuthority")                                     // 设置用户权限

		userGroup.POST("/changePassword") // 用户修改密码
		userGroup.POST("/resetPassword")  // 管理员重置用户密码

		userGroup.GET("/getUserInfo")  // 获取用户信息
		userGroup.POST("/getUserList") // 管理员分页获取用户列表
	}
}
