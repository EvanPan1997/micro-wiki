package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"micro-wiki/global"
	"micro-wiki/model/system"
	"os"
)

func Gorm() *gorm.DB {
	//switch global.MW_CONFIG.Server.DbType {
	//case "mysql":
	//	return GormMysql()
	//default:
	//	return GormMysql()
	//}
	return GormMysql()
}

func Dsn() string {
	m := global.MW_CONFIG.Mysql
	if m.Config == "" {
		m.Config = "charset=utf8mb4&parseTime=True&loc=Local&parseTime=true"
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Address + ":" + m.Port + ")/" + m.DbName + "?" + m.Config
	//log.Println("dsn=", dsn)
	return dsn
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		system.WikiUser{},
	)
	if err != nil {
		global.MW_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.MW_LOG.Info("register table success")
}
