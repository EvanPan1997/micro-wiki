package main

import (
	"github.com/gin-gonic/gin"
	"micro-wiki/router"
)

func main() {
	r := gin.Default()
	// 初始化路由
	router.InitRoute(r)
	_ = r.Run(":9090")
}
