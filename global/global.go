package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"micro-wiki/config"
)

var (
	MW_CONFIG config.System
	MW_VP     *viper.Viper
	MW_LOG    *zap.Logger
	MW_DB     *gorm.DB
)
