package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/admin/cmd/middleware"
	"sgblog-go/app/admin/cmd/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()

	// cros配置
	Router.Use(middleware.Cors())
	adminRouter := router.RouterGroupApp.Admin

	PublicGroup := Router.Group(global.SG_BLOG_COFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		adminRouter.InitLoginRouter(PublicGroup)
		adminRouter.InitArticleRouter(PublicGroup)
		adminRouter.InitCategoryRouter(PublicGroup)
		adminRouter.InitLinkRouter(PublicGroup)
		adminRouter.InitTagRouter(PublicGroup)
		adminRouter.InitUserRouter(PublicGroup)
		adminRouter.InitMenuRouter(PublicGroup)
		adminRouter.InitRoleRouter(PublicGroup)
	}
	PrivateGroup := Router.Group(global.SG_BLOG_COFIG.System.RouterPrefix)
	{
		PrivateGroup.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, "hello world")
		})
		adminRouter.InitLoginPrivateRouter(PrivateGroup)
		adminRouter.InitArticlePrivateRouter(PrivateGroup)
		adminRouter.InitCategoryPrivateRouter(PrivateGroup)
		adminRouter.InitLinkPrivateRouter(PrivateGroup)
		adminRouter.InitTagPrivateRouter(PrivateGroup)
		adminRouter.InitUserPrivateRouter(PrivateGroup)
		adminRouter.InitMenuPrivateRouter(PrivateGroup)
		adminRouter.InitRolePrivateRouter(PrivateGroup)
	}
	global.SG_BLOG_LOG.Info("router register success")
	return Router
}
