package admin

import (
	"sgblog-go/app/admin/cmd/service"
)

type ApiGroup struct {
	AdminLoginApi
}

var (
	blogCategoryService = service.ServiceGroupApp.BlogServiceGroup.CategoryService
	blogArticleService  = service.ServiceGroupApp.BlogServiceGroup.ArticleService
	blogLoginService    = service.ServiceGroupApp.BlogServiceGroup.LoginService
	blogUserService     = service.ServiceGroupApp.BlogServiceGroup.UserService
	blogCommentService  = service.ServiceGroupApp.BlogServiceGroup.CommentService
	blogLinkService     = service.ServiceGroupApp.BlogServiceGroup.LinkService
	menuService         = service.ServiceGroupApp.BlogServiceGroup.MenuService
	roleService         = service.ServiceGroupApp.BlogServiceGroup.RoleService
)
