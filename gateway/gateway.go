package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"micro-wiki/util"
)

func main() {
	gin.SetMode(util.AppMode)
	// 初始化路由
	r := InitRouter()

	err := r.Run(":" + util.HttpPort)
	if err != nil {
		log.Fatalf("Gin启动失败, 错误原因为:\n%s", err)
	}
}
