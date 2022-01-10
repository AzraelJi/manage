package main

import (
	"zga-client-manage/core"
	"zga-client-manage/global"
	"zga-client-manage/initialize"
)

//go:generate go mod tidy
//go:generate go mod download

// @title Swagger ZGA CLIENT MANAGE
// @version 0.0.1
// @description zga client manage
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core.Viper()           // 初始化Viper
	global.GVA_LOG = core.Zap()            // 初始化zap日志库
	global.GVA_DB = initialize.GormMysql() // gorm连接数据库
	initialize.Timer()
	if global.GVA_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
