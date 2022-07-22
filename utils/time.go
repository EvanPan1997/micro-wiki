package utils

import (
	"log"
	"time"
)

var (
	tpl = "20060102150405"
)

// 时间类型转字符串 yyyyMMddHHmmSS
func TimeToString(t time.Time) string {
	return t.Format(tpl)
}

// 字符串转时间类型
func StringToTime(s string) time.Time {
	t, err := time.Parse(tpl, s)
	if err != nil {
		log.Println(err)
	}
	return t
}
