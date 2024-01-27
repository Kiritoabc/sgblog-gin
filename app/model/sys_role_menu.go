package model

type SysRoleMenu struct {
	RoleId int64 `gorm:"column:role_id;type:BIGINT(20);NOT NULL"`
	MenuId int64 `gorm:"column:menu_id;type:BIGINT(20);NOT NULL"`
}
