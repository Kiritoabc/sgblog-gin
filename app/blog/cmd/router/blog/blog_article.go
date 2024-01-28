package blog

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/blog/cmd/api/v1"
)

type ArticleRouter struct{}

func (s ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	articleRouter := Router.Group("article")
	articleApi := v1.ApiGroupApp.BlogApiGroup.BlogArticleApi
	{
		articleRouter.GET("/hotArticleList", articleApi.HotArticleList)
		articleRouter.GET("/articleList", articleApi.ArticleList)
		articleRouter.GET("/:id", articleApi.GetArticleDetail)
	}
}
