package admin

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/admin/cmd/api/v1"
	"sgblog-go/app/admin/cmd/middleware"
)

type MenuRouter struct {
}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("/system/menu")
	menuApi := v1.ApiGroupApp.BlogApiGroup.AdminMenuApi
	{
		_ = menuRouter
		_ = menuApi
	}
}

func (s *MenuRouter) InitMenuPrivateRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("/system/menu").Use(middleware.JwtAuth())
	menuApi := v1.ApiGroupApp.BlogApiGroup.AdminMenuApi
	{
		_ = menuRouter
		_ = menuApi
		menuRouter.GET("/treeselect", menuApi.TreeSelect)
		menuRouter.GET("/roleMenuTreeselect/:roleId", menuApi.RoleMenuTreeSelect)
		menuRouter.GET("/list", menuApi.List)
		menuRouter.POST("/add", menuApi.Add)
		menuRouter.GET("/:menuId", menuApi.GetInfo)
		menuRouter.PUT("/edit", menuApi.Edit)
		menuRouter.DELETE("/:menuId", menuApi.Remove)
	}
}
