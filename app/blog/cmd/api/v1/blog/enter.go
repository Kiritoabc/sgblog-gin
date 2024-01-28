package blog

import (
	"sgblog-go/app/service"
)

type ApiGroup struct {
	BlogCategoryApi
	BlogArticleApi
	BlogLoginApi
}

var (
	blogCategoryService = service.ServiceGroupApp.BlogServiceGroup.CategoryService
	blogArticleService  = service.ServiceGroupApp.BlogServiceGroup.ArticleService
	blogLoginService    = service.ServiceGroupApp.BlogServiceGroup.LoginService
)
