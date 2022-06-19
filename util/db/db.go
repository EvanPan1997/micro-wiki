package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	cfg "micro-wiki/util"
)

func GetDbConnection() (*gorm.DB, error) {
	dsn := cfg.DbUser + ":" + cfg.DbPassword + "@tcp(" + cfg.DbHost + ":" + cfg.DbPort + ")/" + cfg.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
