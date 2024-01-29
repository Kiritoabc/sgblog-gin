package blog

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/blog/cmd/api/v1"
	"sgblog-go/app/blog/cmd/middleware"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userApi := v1.ApiGroupApp.BlogApiGroup.BlogUserApi

	privateRouter := Router.Group("/user")
	privateRouter.Use(middleware.JwtAuth())
	{
		privateRouter.GET("/user/userInfo", userApi.UserInfo)
		privateRouter.PUT("/user/userInfo", userApi.UpdateUserInfo)
	}
}
