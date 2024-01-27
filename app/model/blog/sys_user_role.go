package blog

type SysUserRole struct {
	UserId int64 `gorm:"column:user_id;type:BIGINT(20);NOT NULL" json:"userId"`
	RoleId int64 `gorm:"column:role_id;type:BIGINT(20);NOT NULL" json:"roleId"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
