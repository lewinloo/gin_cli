package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	DBType string
	DBDsn  string

	AppMode string
	AppPort string

	LogConf LogConfig
	JwtConf JwtConfig

	FilePath = "config.ini"
)

func Init() {
	file, err := ini.Load(FilePath)

	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径")
	}

	LoadApp(file)
	LoadDatabase(file)
	LogConf.Load(file)
	JwtConf.Load(file)
}
