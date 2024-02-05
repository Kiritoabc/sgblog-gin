package main

import (
	"sgblog-go/app/blog/cmd/core"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/blog/cmd/initialize"

	"go.uber.org/zap"
)

var configFile = "app/admin/cmd/etc/admin.yaml"

func main() {
	global.SG_BLOG_VP = core.Viper(configFile) // 初始化Viper
	global.SG_BLOG_LOG = core.Zap()
	zap.ReplaceGlobals(global.SG_BLOG_LOG)
	global.SG_BLOG_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.SG_BLOG_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.SG_BLOG_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
