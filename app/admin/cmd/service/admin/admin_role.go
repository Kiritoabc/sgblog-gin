package admin

import (
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
)

type RoleService struct{}

func (s *RoleService) SelectRoleKeyByUserId(userId int64) ([]string, error) {
	//判断是否是管理员 如果是返回集合中只需要有admin
	if userId == 1 {
		return []string{"admin"}, nil
	}
	//否则查询用户所具有的角色信息
	var roleKeys []string
	err := global.SG_BLOG_DB.Model(&blog.SysRoleMenu{}).
		Joins("left join sys_role on sys_role_menu.role_id = sys_role.role_id").
		Where("sys_role_menu.user_id = ? and sys_role_menu = ? and sys_role_menu = ?", userId, 0, 0).
		Pluck("sys_role.role_key", &roleKeys).Error
	if err != nil {
		return nil, err
	}
	return roleKeys, nil
}
