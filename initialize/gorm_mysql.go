package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"micro-wiki/global"
	"os"
	"time"
)

func GormMysql() *gorm.DB {
	m := global.MW_CONFIG.Mysql
	if m.DbName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN: Dsn(),
		//DefaultStringSize:         512,
		SkipInitializeWithVersion:     false,
		DontSupportNullAsDefaultValue: false,
	}
	//dsn := Dsn()
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		log.Println(err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}

func Config() *gorm.Config {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	config.Logger = _default.LogMode(logger.Info)
	return config
}
