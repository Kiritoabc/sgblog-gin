package model

type SysUserRole struct {
	UserId int64 `gorm:"column:user_id;type:BIGINT(20);NOT NULL"`
	RoleId int64 `gorm:"column:role_id;type:BIGINT(20);NOT NULL"`
}
