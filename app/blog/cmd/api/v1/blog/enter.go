package blog

import "sgblog-go/app/blog/cmd/service"

type ApiGroup struct {
	BlogCategoryApi
}

var (
	blogCategoryService = service.ServiceGroupApp.BlogServiceGroup.CategoryService
)
