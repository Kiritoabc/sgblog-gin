package admin

import (
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/common/constants"
)

type MenuService struct{}

// SelectPermsByUserId 根据用户id查询权限信息
func (s *MenuService) SelectPermsByUserId(userId int64, isAdmin bool) ([]string, error) {
	// 1. 如果是管理员，返回所有的权限
	if isAdmin {
		var menus []*blog.SysMenu
		err := global.SG_BLOG_DB.Model(&blog.SysMenu{}).
			Where("menu_type in (?)", []string{constants.Menu, constants.Button}).
			Where("status = ?", constants.STATUS_NORMAL).
			Find(&menus).Error
		if err != nil {
			return nil, err
		}
		var perms []string
		for _, menu := range menus {
			perms = append(perms, menu.Perms)
		}
		return perms, nil
	}
	// 2.非管理员查询权限返回
	return s.selectPermsByUserId(userId)
}

func (s *MenuService) selectPermsByUserId(userId int64) ([]string, error) {
	var permission []string
	err := global.SG_BLOG_DB.Model(&blog.SysUserRole{}).
		Joins("LEFT JOIN sys_role_menu ON sys_user_role.role_id = sys_role_menu.role_id").
		Joins("LEFT JOIN sys_menu ON sys_menu.id = sys_role_menu.menu_id").
		Where("sys_user_role.user_id = ?", userId).
		Where("sys_menu.menu_type IN (?) AND sys_menu.status = ? AND sys_menu.del_flag = ?", []string{"C", "F"}, 0, 0).
		Pluck("DISTINCT sys_menu.perms", &permission).Error
	if err != nil {
		return nil, err
	}
	return permission, nil
}

//

func (s *MenuService) selectRouterMenuTreeByUserId(userId int64) {

}
