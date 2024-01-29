package blog

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/blog/cmd/api/v1"
	"sgblog-go/app/blog/cmd/middleware"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.JwtAuth())
	userApi := v1.ApiGroupApp.BlogApiGroup.BlogUserApi
	{
		userRouter.GET("/userInfo", userApi.UserInfo)
	}
}
