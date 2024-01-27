package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/model/blog"
)

func Gorm() *gorm.DB {
	switch global.SG_BLOG_COFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables() {
	db := global.SG_BLOG_DB
	err := db.AutoMigrate(
		// 系统模块表
		blog.SgArticle{},
		blog.SgArticleTag{},
		blog.SgCategory{},
		blog.SgComment{},
		blog.SgLink{},
		blog.SgTag{},
		blog.SysMenu{},
		blog.SysRole{},
		blog.SysRoleMenu{},
		blog.SysUser{},
		blog.SysUserRole{},
	)
	if err != nil {
		global.SG_BLOG_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.SG_BLOG_LOG.Info("register table success")
}
