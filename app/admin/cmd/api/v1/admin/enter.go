package admin

import (
	"sgblog-go/app/admin/cmd/service"
)

type ApiGroup struct {
	AdminLoginApi
	AdminArticleApi
	AdminCategoryApi
	AdminLinkApi
	AdminTagApi
	AdminUserApi
	AdminMenuApi
}

var (
	blogCategoryService = service.ServiceGroupApp.BlogServiceGroup.CategoryService
	blogArticleService  = service.ServiceGroupApp.BlogServiceGroup.ArticleService
	blogLoginService    = service.ServiceGroupApp.BlogServiceGroup.LoginService
	blogUserService     = service.ServiceGroupApp.BlogServiceGroup.UserService
	blogCommentService  = service.ServiceGroupApp.BlogServiceGroup.CommentService
	blogLinkService     = service.ServiceGroupApp.BlogServiceGroup.LinkService
	blogTagService      = service.ServiceGroupApp.BlogServiceGroup.TagService
	menuService         = service.ServiceGroupApp.BlogServiceGroup.MenuService
	roleService         = service.ServiceGroupApp.BlogServiceGroup.RoleService
)
