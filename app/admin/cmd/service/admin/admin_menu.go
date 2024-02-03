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

func (s *MenuService) SelectRouterMenuTreeByUserId(userId int64,
	isAdmin bool) ([]*blog.SysMenu, error) {
	var menus []*blog.SysMenu
	if isAdmin {
		menus, _ = s.selectAllRouterMenu()
	} else {
		menus, _ = s.selectRouterMenuTreeByUserId(userId)
	}
	menuTree, err := buildMenuTree(menus, 0)
	if err != nil {
		return nil, err
	}
	return menuTree, nil
}

func (s *MenuService) selectAllRouterMenu() ([]*blog.SysMenu, error) {
	var menus []*blog.SysMenu
	db := global.SG_BLOG_DB
	err := db.Table("sys_menu m").
		Select("DISTINCT m.id, m.parent_id, m.menu_name, m.path, m.component, m.visible, "+
			"m.status, IFNULL(m.perms, '') as perms, m.is_frame, m.menu_type, m.icon, m.order_num, m.create_time").
		Where("m.menu_type IN (?) AND m.status = ? AND m.del_flag = ?", []string{"C", "M"}, 0, 0).
		Order("m.parent_id, m.order_num").
		Find(&menus).Error

	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (s *MenuService) selectRouterMenuTreeByUserId(userId int64) ([]*blog.SysMenu, error) {
	var menus []*blog.SysMenu
	db := global.SG_BLOG_DB
	err := db.Table("sys_user_role ur").Joins("LEFT JOIN sys_role_menu rm ON ur.role_id = rm.role_id").
		Joins("LEFT JOIN sys_menu m ON m.id = rm.menu_id").
		Select("DISTINCT m.id, m.parent_id, m.menu_name, m.path, m.component, m.visible, m.status, IFNULL(m.perms, '') as perms, "+
			"m.is_frame, m.menu_type, m.icon, m.order_num, m.create_time").
		Where("ur.user_id = ? AND m.menu_type IN (?,?) AND m.status = ? AND m.del_flag = ?", userId, "C", "M", 0, 0).
		Order("m.parent_id, m.order_num").
		Find(&menus).Error

	if err != nil {
		return nil, err
	}

	return menus, nil
}

func buildMenuTree(menus []*blog.SysMenu, parentId int64) ([]*blog.SysMenu, error) {
	var menuTree []*blog.SysMenu
	for _, menu := range menus {
		if menu.ParentId == parentId {
			menu.Children, _ = getChildren(menu, menus)
			menuTree = append(menuTree, menu)
		}
	}
	return menuTree, nil
}

func getChildren(menu *blog.SysMenu, menus []*blog.SysMenu) ([]*blog.SysMenu, error) {
	var children []*blog.SysMenu
	for _, m := range menus {
		if m.ParentId == menu.Id {
			mCopy := m
			mCopy.Children, _ = getChildren(mCopy, menus)
			children = append(children, mCopy)
		}
	}
	return children, nil
}
