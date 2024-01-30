package blog

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/blog/cmd/api/v1"
	"sgblog-go/app/blog/cmd/middleware"
)

type CommentRouter struct{}

func (s *CommentRouter) InitCommentRouter(Router *gin.RouterGroup) {
	commentRouter := Router.Group("comment")
	commentApi := v1.ApiGroupApp.BlogApiGroup.BlogCommentApi
	{
		commentRouter.GET("/commentList", commentApi.CommentList)         // 获取分类列表
		commentRouter.GET("/linkCommentList", commentApi.LinkCommentList) // 获取友联评论
	}
}

func (s *CommentRouter) InitCommentPrivateRouter(Router *gin.RouterGroup) {
	commentRouter := Router.Group("comment").Use(middleware.JwtAuth())
	commentApi := v1.ApiGroupApp.BlogApiGroup.BlogCommentApi
	{
		commentRouter.POST("/", commentApi.AddComment) // 添加评论
	}
}
