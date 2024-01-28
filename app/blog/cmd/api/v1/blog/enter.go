package blog

import (
	"sgblog-go/app/service"
)

type ApiGroup struct {
	BlogCategoryApi
	BlogArticleApi
}

var (
	blogCategoryService = service.ServiceGroupApp.BlogServiceGroup.CategoryService
	blogArticleService  = service.ServiceGroupApp.BlogServiceGroup.ArticleService
)
