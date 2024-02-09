package admin

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/admin/cmd/api/v1"
	"sgblog-go/app/admin/cmd/middleware"
)

type RoleRouter struct {
}

func (s *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleRouter := Router.Group("/system/role")
	roleApi := v1.ApiGroupApp.BlogApiGroup.AdminRoleApi
	{
		_ = roleRouter
		_ = roleApi
	}
}

func (s *RoleRouter) InitRolePrivateRouter(Router *gin.RouterGroup) {
	roleRouter := Router.Group("/system/role").Use(middleware.JwtAuth())
	roleApi := v1.ApiGroupApp.BlogApiGroup.AdminRoleApi
	{
		_ = roleRouter
		_ = roleApi
		roleRouter.GET("/listAllRole", roleApi.ListAllRole)
		roleRouter.GET("/:roleId", roleApi.GetInfo)
		roleRouter.POST("", roleApi.Add)
		roleRouter.PUT("", roleApi.Edit)
		roleRouter.DELETE("/:id", roleApi.Remove)
		roleRouter.POST("/changeStatus", roleApi.ChangeStatus)
		roleRouter.GET("/list", roleApi.List)
	}
}
