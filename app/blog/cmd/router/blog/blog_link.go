package blog

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/blog/cmd/api/v1"
)

type LinkRouter struct{}

func (s *LinkRouter) InitLinkRouter(Router *gin.RouterGroup) {
	linkRouter := Router.Group("link")
	linkApi := v1.ApiGroupApp.BlogApiGroup.BlogLinkApi
	{
		linkRouter.GET("/getAllLink", linkApi.GetAllLink) // 获取所有友链
	}
}
