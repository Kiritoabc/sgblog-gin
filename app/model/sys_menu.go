package model

import "time"

type SysMenu struct {
	Id         int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	MenuName   string    `gorm:"column:menu_name;type:VARCHAR(50);NOT NULL"`
	ParentId   int64     `gorm:"column:parent_id;type:BIGINT(20);"`
	OrderNum   int32     `gorm:"column:order_num;type:INT(4);"`
	Path       string    `gorm:"column:path;type:VARCHAR(200);"`
	Component  string    `gorm:"column:component;type:VARCHAR(255);"`
	IsFrame    int32     `gorm:"column:is_frame;type:INT(1);"`
	MenuType   string    `gorm:"column:menu_type;type:CHAR(1);"`
	Visible    string    `gorm:"column:visible;type:CHAR(1);"`
	Status     string    `gorm:"column:status;type:CHAR(1);"`
	Perms      string    `gorm:"column:perms;type:VARCHAR(100);"`
	Icon       string    `gorm:"column:icon;type:VARCHAR(100);"`
	CreateBy   int64     `gorm:"column:create_by;type:BIGINT(20);"`
	CreateTime time.Time `gorm:"column:create_time;type:DATETIME;"`
	UpdateBy   int64     `gorm:"column:update_by;type:BIGINT(20);"`
	UpdateTime time.Time `gorm:"column:update_time;type:DATETIME;"`
	Remark     string    `gorm:"column:remark;type:VARCHAR(500);"`
	DelFlag    string    `gorm:"column:del_flag;type:CHAR(1);"`
}
