package util

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode  string
	HttpPort string

	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln("配置文件读取错误, 请检查文件路径:", err)
	}
	LoadServer(file)
}

func LoadServer(file *ini.File) {
	// Server
	AppMode = file.Section("Server").Key("RunMode").MustString("debug")
	HttpPort = file.Section("Server").Key("HttpPort").MustString("8080")

	// database
	DbUser = file.Section("Database").Key("DbUser").MustString("root")
	DbPassword = file.Section("Database").Key("DbPassword").MustString("root")
	DbHost = file.Section("Database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("Database").Key("DbPort").MustString("3306")
	DbName = file.Section("Database").Key("DbName").MustString("micro_wiki")
}
