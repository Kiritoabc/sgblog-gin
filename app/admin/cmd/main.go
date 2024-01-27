package main

import (
	"go.uber.org/zap"
	"sgblog-go/app/admin/cmd/core"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/admin/cmd/initialize"
)

var configPath = "app/admin/cmd/etc/admin.yaml"

func main() {
	global.SG_BLOG_VP = core.Viper(configPath) // 初始化Viper
	global.SG_BLOG_LOG = core.Zap()
	zap.ReplaceGlobals(global.SG_BLOG_LOG)
	global.SG_BLOG_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()
	if global.SG_BLOG_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.SG_BLOG_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
