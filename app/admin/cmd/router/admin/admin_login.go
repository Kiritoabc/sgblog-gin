package admin

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/admin/cmd/api/v1"
	"sgblog-go/app/admin/cmd/middleware"
)

type LoginRouter struct {
}

func (s LoginRouter) InitLoginROuter(Router *gin.RouterGroup) {
	loginRouter := Router.Group("user")
	loginApi := v1.ApiGroupApp.BlogApiGroup.AdminLoginApi
	{
		loginRouter.POST("/login", loginApi.Login)                             // 管理员登录
		loginRouter.Use(middleware.JwtAuth()).POST("/logout", loginApi.Logout) // 管理员退出
	}
}
