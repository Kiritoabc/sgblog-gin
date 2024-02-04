package admin

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/admin/cmd/api/v1"
	"sgblog-go/app/admin/cmd/middleware"
)

type LinkRouter struct {
}

func (s *LinkRouter) InitLinkRouter(Router *gin.RouterGroup) {
	linkRouter := Router.Group("link")
	linkApi := v1.ApiGroupApp.BlogApiGroup.AdminLinkApi
	{
		_ = linkRouter
		_ = linkApi
	}
}

func (s *LinkRouter) InitLinkPrivateRouter(Router *gin.RouterGroup) {
	linkRouter := Router.Group("/content/link").Use(middleware.JwtAuth())
	linkApi := v1.ApiGroupApp.BlogApiGroup.AdminLinkApi
	{
		linkRouter.GET("/list", linkApi.List)                         // get link list
		linkRouter.POST("", linkApi.Add)                              // create a link record
		linkRouter.DELETE("/:id", linkApi.Delete)                     // remove a record by id
		linkRouter.PUT("", linkApi.Edit)                              // update link
		linkRouter.PUT("/changeLinkStatus", linkApi.ChangeLinkStatus) // change link status
		linkRouter.GET("/:id", linkApi.GetInfo)                       // get link info by id
	}
}
