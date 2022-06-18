package util

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode  string
	HttpPort string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln("配置文件读取错误, 请检查文件路径:", err)
	}
	LoadServer(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("Server").Key("RunMode").MustString("debug")
	HttpPort = file.Section("Server").Key("HttpPort").MustString("8080")
}
