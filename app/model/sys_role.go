package model

import "time"

type SysRole struct {
	Id         int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	RoleName   string    `gorm:"column:role_name;type:VARCHAR(30);NOT NULL"`
	RoleKey    string    `gorm:"column:role_key;type:VARCHAR(100);NOT NULL"`
	RoleSort   int32     `gorm:"column:role_sort;type:INT(4);NOT NULL"`
	Status     string    `gorm:"column:status;type:CHAR(1);NOT NULL"`
	DelFlag    string    `gorm:"column:del_flag;type:CHAR(1);"`
	CreateBy   int64     `gorm:"column:create_by;type:BIGINT(20);"`
	CreateTime time.Time `gorm:"column:create_time;type:DATETIME;"`
	UpdateBy   int64     `gorm:"column:update_by;type:BIGINT(20);"`
	UpdateTime time.Time `gorm:"column:update_time;type:DATETIME;"`
	Remark     string    `gorm:"column:remark;type:VARCHAR(500);"`
}
