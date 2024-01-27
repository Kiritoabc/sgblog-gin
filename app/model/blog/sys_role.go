package blog

import "time"

type SysRole struct {
	Id         int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL" json:"id"`
	RoleName   string    `gorm:"column:role_name;type:VARCHAR(30);NOT NULL" json:"roleName"`
	RoleKey    string    `gorm:"column:role_key;type:VARCHAR(100);NOT NULL" json:"roleKey"`
	RoleSort   int32     `gorm:"column:role_sort;type:INT(4);NOT NULL" json:"roleSort"`
	Status     string    `gorm:"column:status;type:CHAR(1);NOT NULL" json:"status"`
	DelFlag    string    `gorm:"column:del_flag;type:CHAR(1);" json:"delFlag"`
	CreateBy   int64     `gorm:"column:create_by;type:BIGINT(20);" json:"createBy"`
	CreateTime time.Time `gorm:"column:create_time;type:DATETIME;" json:"createTime"`
	UpdateBy   int64     `gorm:"column:update_by;type:BIGINT(20);" json:"updateBy"`
	UpdateTime time.Time `gorm:"column:update_time;type:DATETIME;" json:"updateTime"`
	Remark     string    `gorm:"column:remark;type:VARCHAR(500);" json:"remark"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
