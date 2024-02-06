package admin

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/admin/cmd/api/v1"
	"sgblog-go/app/admin/cmd/middleware"
)

type UserRouter struct {
}

func (s UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("/system/user")
	userApi := v1.ApiGroupApp.BlogApiGroup.AdminUserApi
	{
		_ = userRouter
		_ = userApi
	}
}

func (s UserRouter) InitUserPrivateRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("/system/user").Use(middleware.JwtAuth())
	userApi := v1.ApiGroupApp.BlogApiGroup.AdminUserApi
	{
		_ = userRouter
		_ = userApi
		userRouter.GET("/list", userApi.List)
		userRouter.POST("", userApi.Add)
		userRouter.GET("/:userId", userApi.GetUserInfoAndRoleIds)
		userRouter.POST("", userApi.Edit)
		userRouter.DELETE("/:userIds", userApi.Remove)
	}
}
