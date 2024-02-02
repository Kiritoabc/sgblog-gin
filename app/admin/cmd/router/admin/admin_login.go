package admin

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/admin/cmd/api/v1"
	"sgblog-go/app/admin/cmd/middleware"
)

type LoginRouter struct {
}

func (s LoginRouter) InitLoginROuter(Router *gin.RouterGroup) {

	loginApi := v1.ApiGroupApp.BlogApiGroup.AdminLoginApi
	{
		Router.POST("/user/login", loginApi.Login) // 管理员登录

	}
}

func (s LoginRouter) InitLoginPrivateRouter(Router *gin.RouterGroup) {
	loginApi := v1.ApiGroupApp.BlogApiGroup.AdminLoginApi
	Router.Use(middleware.JwtAuth())
	{
		Router.POST("/user/logout", loginApi.Logout) // 管理员退出
		Router.GET("/getInfo", loginApi.GetInfo)     // 获取信息
	}
}
