package initialize

import (
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// Author [SliverHorn](https://github.com/SliverHorn)
func (w *writer) Printf(message string, data ...interface{}) {
	//flag := true
	//if flag {
	//	global.MW_LOG.Info(fmt.Sprintf(message+"\n", data...))
	//} else {
	//	w.Writer.Printf(message, data...)
	//}

	//global.MW_LOG.Info(fmt.Sprintf(message+"\n", data...))
	w.Writer.Printf(message, data...)
}
