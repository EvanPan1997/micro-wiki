package main

import (
	"github.com/gin-gonic/gin"
	"micro-wiki/core"
	"micro-wiki/global"
	"micro-wiki/initialize"
	"micro-wiki/router"
)

func main() {
	global.MW_VP = core.Viper()
	gin.SetMode(gin.DebugMode)
	global.MW_DB = initialize.Gorm()
	if global.MW_DB != nil {
		//initialize.RegisterTables(global.MW_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.MW_DB.DB()
		defer db.Close()
	}
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	// 初始化路由
	router.InitRoute(r)
	_ = r.Run(":9090")
}
