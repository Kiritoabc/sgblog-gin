package model

import "time"

type SgComment struct {
	Id              int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	Type            string    `gorm:"column:type;type:CHAR(1);"`
	ArticleId       int64     `gorm:"column:article_id;type:BIGINT(20);"`
	RootId          int64     `gorm:"column:root_id;type:BIGINT(20);"`
	Content         string    `gorm:"column:content;type:VARCHAR(512);"`
	ToCommentUserId int64     `gorm:"column:to_comment_user_id;type:BIGINT(20);"`
	ToCommentId     int64     `gorm:"column:to_comment_id;type:BIGINT(20);"`
	CreateBy        int64     `gorm:"column:create_by;type:BIGINT(20);"`
	CreateTime      time.Time `gorm:"column:create_time;type:DATETIME;"`
	UpdateBy        int64     `gorm:"column:update_by;type:BIGINT(20);"`
	UpdateTime      time.Time `gorm:"column:update_time;type:DATETIME;"`
	DelFlag         int32     `gorm:"column:del_flag;type:INT(1);"`
}
