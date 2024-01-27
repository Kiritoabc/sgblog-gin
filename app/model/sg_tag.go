package model

import "time"

type SgTag struct {
	Id         int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	Name       string    `gorm:"column:name;type:VARCHAR(128);"`
	CreateBy   int64     `gorm:"column:create_by;type:BIGINT(20);"`
	CreateTime time.Time `gorm:"column:create_time;type:DATETIME;"`
	UpdateBy   int64     `gorm:"column:update_by;type:BIGINT(20);"`
	UpdateTime time.Time `gorm:"column:update_time;type:DATETIME;"`
	DelFlag    int32     `gorm:"column:del_flag;type:INT(1);"`
	Remark     string    `gorm:"column:remark;type:VARCHAR(500);"`
}
