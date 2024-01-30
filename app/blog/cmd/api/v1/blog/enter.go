package blog

import (
	"sgblog-go/app/service"
)

type ApiGroup struct {
	BlogCategoryApi
	BlogArticleApi
	BlogLoginApi
	BlogUserApi
	BlogCommentApi
	BlogLinkApi
}

var (
	blogCategoryService = service.ServiceGroupApp.BlogServiceGroup.CategoryService
	blogArticleService  = service.ServiceGroupApp.BlogServiceGroup.ArticleService
	blogLoginService    = service.ServiceGroupApp.BlogServiceGroup.LoginService
	blogUserService     = service.ServiceGroupApp.BlogServiceGroup.UserService
	blogCommentService  = service.ServiceGroupApp.BlogServiceGroup.CommentService
	blogLinkService     = service.ServiceGroupApp.BlogServiceGroup.LinkService
)
