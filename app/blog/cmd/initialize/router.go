package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "sgblog-go/app/blog/cmd/api/v1"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/blog/cmd/middleware"
	"sgblog-go/app/blog/cmd/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.New()

	Router.Use(middleware.Cors())

	blogRouter := router.RouterGroupApp.Blog
	// 注册路由
	Router.POST("/user/register", v1.ApiGroupApp.BlogApiGroup.Register)

	PublicGroup := Router.Group(global.SG_BLOG_COFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		blogRouter.InitCategoryRouter(PublicGroup)
		blogRouter.InitArticleRouter(PublicGroup)
		blogRouter.InitCommentRouter(PublicGroup)
		blogRouter.InitLinkRouter(PublicGroup)
	}
	PrivateGroup := Router.Group(global.SG_BLOG_COFIG.System.RouterPrefix)
	{
		PrivateGroup.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, "hello world")
		})
		_ = blogRouter
		blogRouter.InitUserRouter(PrivateGroup)
		blogRouter.InitLoginRouter(PrivateGroup)
		blogRouter.InitCommentPrivateRouter(PrivateGroup)
	}
	global.SG_BLOG_LOG.Info("router register success")
	return Router
}
