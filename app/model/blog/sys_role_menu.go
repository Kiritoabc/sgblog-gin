package blog

type SysRoleMenu struct {
	RoleId int64 `gorm:"column:role_id;type:BIGINT(20);NOT NULL" json:"roleId"`
	MenuId int64 `gorm:"column:menu_id;type:BIGINT(20);NOT NULL" json:"menuId"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
