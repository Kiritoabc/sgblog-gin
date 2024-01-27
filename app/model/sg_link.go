package model

import "time"

type SgLink struct {
	Id          int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	Name        string    `gorm:"column:name;type:VARCHAR(256);"`
	Logo        string    `gorm:"column:logo;type:VARCHAR(256);"`
	Description string    `gorm:"column:description;type:VARCHAR(512);"`
	Address     string    `gorm:"column:address;type:VARCHAR(128);"`
	Status      string    `gorm:"column:status;type:CHAR(1);"`
	CreateBy    int64     `gorm:"column:create_by;type:BIGINT(20);"`
	CreateTime  time.Time `gorm:"column:create_time;type:DATETIME;"`
	UpdateBy    int64     `gorm:"column:update_by;type:BIGINT(20);"`
	UpdateTime  time.Time `gorm:"column:update_time;type:DATETIME;"`
	DelFlag     int32     `gorm:"column:del_flag;type:INT(1);"`
}
