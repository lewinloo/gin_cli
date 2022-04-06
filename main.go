package main

import (
	"gin_cli/app"
	"gin_cli/config"
	"gin_cli/global"
	"gin_cli/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	// 初始化配置文件
	config.Init()

	global.DB = initialize.Gorm()

	if global.DB != nil {
		// 生成表到数据库
		initialize.GenerateTables(global.DB)
		db, _ := global.DB.DB()
		defer db.Close()
	}

	if logger, err := initialize.InitLogger(&config.LogConf); err != nil {
		panic(err)
	} else {
		global.LOG = logger
	}

	app.Run(config.AppPort)
}
