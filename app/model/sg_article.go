package model

import "time"

type SgArticle struct {
	Id         int64     `gorm:"column:id;type:BIGINT(200);AUTO_INCREMENT;NOT NULL"`
	Title      string    `gorm:"column:title;type:VARCHAR(256);"`
	Content    string    `gorm:"column:content;type:LONGTEXT;"`
	Summary    string    `gorm:"column:summary;type:VARCHAR(1024);"`
	CategoryId int64     `gorm:"column:category_id;type:BIGINT(20);"`
	Thumbnail  string    `gorm:"column:thumbnail;type:VARCHAR(256);"`
	IsTop      string    `gorm:"column:is_top;type:CHAR(1);"`
	Status     string    `gorm:"column:status;type:CHAR(1);"`
	ViewCount  int64     `gorm:"column:view_count;type:BIGINT(200);"`
	IsComment  string    `gorm:"column:is_comment;type:CHAR(1);"`
	CreateBy   int64     `gorm:"column:create_by;type:BIGINT(20);"`
	CreateTime time.Time `gorm:"column:create_time;type:DATETIME;"`
	UpdateBy   int64     `gorm:"column:update_by;type:BIGINT(20);"`
	UpdateTime time.Time `gorm:"column:update_time;type:DATETIME;"`
	DelFlag    int32     `gorm:"column:del_flag;type:INT(1);"`
}
