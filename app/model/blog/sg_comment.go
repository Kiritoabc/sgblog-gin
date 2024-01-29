package blog

import "time"

type SgComment struct {
	Id              int64     `gorm:"column:id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL" json:"id"`
	Type            string    `gorm:"column:type;type:CHAR(1);" json:"type"`
	ArticleId       int64     `gorm:"column:article_id;type:BIGINT(20);" json:"articleId"`
	RootId          int64     `gorm:"column:root_id;type:BIGINT(20);" json:"rootId"`
	Content         string    `gorm:"column:content;type:VARCHAR(512);" json:"content"`
	ToCommentUserId int64     `gorm:"column:to_comment_user_id;type:BIGINT(20);" json:"toCommentUserId"`
	ToCommentId     int64     `gorm:"column:to_comment_id;type:BIGINT(20);" json:"toCommentId"`
	CreateBy        int64     `gorm:"column:create_by;type:BIGINT(20);" json:"createBy"`
	CreateTime      time.Time `gorm:"column:create_time;type:DATETIME;default:current_timestamp" json:"createTime"`
	UpdateBy        int64     `gorm:"column:update_by;type:BIGINT(20);" json:"updateBy"`
	UpdateTime      time.Time `gorm:"column:update_time;type:DATETIME;default:current_timestamp on update current_timestamp" json:"updateTime"`
	DelFlag         int32     `gorm:"column:del_flag;type:INT(1);" json:"delFlag"`
}

func (SgComment) TableName() string {
	return "sg_comment"
}
