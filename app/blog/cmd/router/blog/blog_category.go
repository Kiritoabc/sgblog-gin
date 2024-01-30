package blog

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/blog/cmd/api/v1"
)

type CategoryRouter struct{}

func (s *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	categoryRouter := Router.Group("category")
	categoryApi := v1.ApiGroupApp.BlogApiGroup.BlogCategoryApi
	{
		categoryRouter.GET("/getCategoryList", categoryApi.GetCategoryList) // 获取分类列表
	}
}
