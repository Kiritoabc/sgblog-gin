package model

import "time"

type SgCategory struct {
	Id          int64     `gorm:"column:id;type:BIGINT(200);AUTO_INCREMENT;NOT NULL"`
	Name        string    `gorm:"column:name;type:VARCHAR(128);"`
	Pid         int64     `gorm:"column:pid;type:BIGINT(200);"`
	Description string    `gorm:"column:description;type:VARCHAR(512);"`
	Status      string    `gorm:"column:status;type:CHAR(1);"`
	CreateBy    int64     `gorm:"column:create_by;type:BIGINT(200);"`
	CreateTime  time.Time `gorm:"column:create_time;type:DATETIME;"`
	UpdateBy    int64     `gorm:"column:update_by;type:BIGINT(200);"`
	UpdateTime  time.Time `gorm:"column:update_time;type:DATETIME;"`
	DelFlag     int32     `gorm:"column:del_flag;type:INT(11);"`
}
