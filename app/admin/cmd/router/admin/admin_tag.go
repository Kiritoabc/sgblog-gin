package admin

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/admin/cmd/api/v1"
	"sgblog-go/app/admin/cmd/middleware"
)

type TagRouter struct {
}

func (s TagRouter) InitTagRouter(Router *gin.RouterGroup) {
	tagRouter := Router.Group("tag")
	tagApi := v1.ApiGroupApp.BlogApiGroup.AdminTagApi
	{
		_ = tagRouter
		_ = tagApi
	}
}

func (s TagRouter) InitTagPrivateRouter(Router *gin.RouterGroup) {
	tagRouter := Router.Group("/content/tag").Use(middleware.JwtAuth())
	tagApi := v1.ApiGroupApp.BlogApiGroup.AdminTagApi
	{
		_ = tagRouter
		_ = tagApi
		tagRouter.GET("/list", tagApi.List)
		tagRouter.POST("", tagApi.Add)          // create a link record
		tagRouter.DELETE("/:id", tagApi.Delete) // remove a record by id
		tagRouter.PUT("", tagApi.Edit)
		tagRouter.GET("/:id", tagApi.GetInfo)
		tagRouter.GET("/listAllTag", tagApi.ListAllTag)
	}
}
