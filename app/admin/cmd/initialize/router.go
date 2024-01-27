package initialize

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/admin/cmd/global"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()

	global.SG_BLOG_LOG.Info("router register success")
	return Router
}
