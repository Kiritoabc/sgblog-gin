package admin

import (
	"github.com/gin-gonic/gin"
	v1 "sgblog-go/app/admin/cmd/api/v1"
	"sgblog-go/app/admin/cmd/middleware"
)

type CategoryRouter struct {
}

func (s *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	categoryRouter := Router.Group("category")
	categoryApi := v1.ApiGroupApp.BlogApiGroup.AdminCategoryApi
	{
		_ = categoryRouter
		_ = categoryApi
	}
}

func (s *CategoryRouter) InitCategoryPrivateRouter(Router *gin.RouterGroup) {
	categoryRouter := Router.Group("/content/category").Use(middleware.JwtAuth())
	categoryApi := v1.ApiGroupApp.BlogApiGroup.AdminCategoryApi
	{
		_ = categoryRouter
		_ = categoryApi
		categoryRouter.GET("/listAllCategory", categoryApi.ListAllCategory) // get all categories
		categoryRouter.PUT("", categoryApi.Edit)                            // edit category
		categoryRouter.DELETE("/:id", categoryApi.Remove)                   // remove category by id
		categoryRouter.POST("", categoryApi.Add)                            // create category
		categoryRouter.GET("/:id", categoryApi.GetInfo)                     // get category info
		categoryRouter.GET("/list", categoryApi.List)                       // get category list
	}
}
