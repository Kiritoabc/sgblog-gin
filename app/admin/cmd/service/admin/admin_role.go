package admin

import (
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/common/constants"
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

func (s *RoleService) SelectRoleAll() ([]*blog.SysRole, error) {
	var roleList []*blog.SysRole
	err := global.SG_BLOG_DB.Model(&blog.SysRole{}).Where("status = ?", constants.Normal).
		Find(&roleList).Error

	if err != nil {
		return nil, err
	}

	return roleList, nil
}

func (s *RoleService) SelectRoleIdByUserId(userId int64) ([]int64, error) {
	var ids []int64

	err := global.SG_BLOG_DB.Table("sys_role r").
		Joins("left join sys_user_role ur on ur.role_id = u.id").
		Where("ur.user_id = ?", userId).Pluck("r.id", &ids).Error

	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (s *RoleService) UpdateRole(role blog.SysRole) error {
	tx := global.SG_BLOG_DB.Begin()

	if err := tx.Model(&blog.SysRole{}).Where("id = ?", role.Id).Updates(&role).Error; err != nil {
		if tx.Rollback().Error != nil {
			return err
		}
		return err
	}
	if err := tx.Where("role_id = ?", role.Id).Delete(&blog.SysRoleMenu{}).Error; err != nil {
		if tx.Rollback().Error != nil {
			return err
		}
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
