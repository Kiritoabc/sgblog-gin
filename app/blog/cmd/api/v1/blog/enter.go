package blog

import (
	"sgblog-go/app/service"
)

type ApiGroup struct {
	BlogCategoryApi
	BlogArticleApi
	BlogLoginApi
	BlogUserApi
}

var (
	blogCategoryService = service.ServiceGroupApp.BlogServiceGroup.CategoryService
	blogArticleService  = service.ServiceGroupApp.BlogServiceGroup.ArticleService
	blogLoginService    = service.ServiceGroupApp.BlogServiceGroup.LoginService
	blogUserService     = service.ServiceGroupApp.BlogServiceGroup.UserService
)
