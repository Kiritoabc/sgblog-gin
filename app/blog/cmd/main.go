package main

import (
	"go.uber.org/zap"
	core2 "sgblog-go/common/core"
	"sgblog-go/common/global"
	initialize2 "sgblog-go/common/initialize"
)

var configFile = "app/blog/cmd/etc/blog.yaml"

func main() {
	global.SG_BLOG_VP = core2.Viper(configFile) // 初始化Viper
	global.SG_BLOG_LOG = core2.Zap()
	zap.ReplaceGlobals(global.SG_BLOG_LOG)
	global.SG_BLOG_DB = initialize2.Gorm() // gorm连接数据库
	initialize2.DBList()
	if global.SG_BLOG_DB != nil {
		initialize2.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.SG_BLOG_DB.DB()
		defer db.Close()
	}
	core2.RunWindowsServer()
}
