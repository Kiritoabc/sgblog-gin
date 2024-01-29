package initialize

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/blog/cmd/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	// TODO:暂时不做跨域处理，后续需要再做处理
	Router.Use(cors.Default())
	blogRouter := router.RouterGroupApp.Blog
	PublicGroup := Router.Group(global.SG_BLOG_COFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		blogRouter.InitCategoryRouter(PublicGroup)
		blogRouter.InitArticleRouter(PublicGroup)
		blogRouter.InitLoginRouter(PublicGroup)
		blogRouter.InitUserRouter(PublicGroup)
	}
	PrivateGroup := Router.Group(global.SG_BLOG_COFIG.System.RouterPrefix)
	{
		PrivateGroup.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, "hello world")
		})
		_ = blogRouter
	}
	global.SG_BLOG_LOG.Info("router register success")
	return Router
}
