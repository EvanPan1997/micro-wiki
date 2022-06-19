package main

import (
	"github.com/gin-gonic/gin"
	"log"
	conf "micro-wiki/util"
)

func main() {
	gin.SetMode(conf.AppMode)
	// 初始化路由
	r := InitRouter()

	err := r.Run(":" + conf.HttpPort)
	if err != nil {
		log.Fatalf("Gin启动失败, 错误原因为:\n%s", err)
	}
}
