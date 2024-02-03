package admin

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/admin/cmd/api/v1"
	"sgblog-go/app/admin/cmd/middleware"
)

type ArticleRouter struct {
}

func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	articleApi := v1.ApiGroupApp.BlogApiGroup.AdminArticleApi
	Router.Group("")
	{
		_ = articleApi
	}
}
func (s *ArticleRouter) InitArticlePrivateRouter(Router *gin.RouterGroup) {
	articleApi := v1.ApiGroupApp.BlogApiGroup.AdminArticleApi
	iRoutes := Router.Group("/content/article").Use(middleware.JwtAuth())
	{
		iRoutes.POST("", articleApi.Add)
		iRoutes.PUT("", articleApi.Edit)
		iRoutes.DELETE("/:id", articleApi.Delete)
		iRoutes.GET("/:id", articleApi.GetInfo)
		iRoutes.GET("/list", articleApi.List)
	}
}
