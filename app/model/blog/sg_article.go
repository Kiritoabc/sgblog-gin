package blog

import "time"
import "gorm.io/plugin/soft_delete"

type SgArticle struct {
	Id         int64                 `gorm:"column:id;type:BIGINT(200);AUTO_INCREMENT;NOT NULL" json:"id"`
	Title      string                `gorm:"column:title;type:VARCHAR(256);" json:"title"`
	Content    string                `gorm:"column:content;type:LONGTEXT;" json:"content"`
	Summary    string                `gorm:"column:summary;type:VARCHAR(1024);" json:"summary"`
	CategoryId int64                 `gorm:"column:category_id;type:BIGINT(20);" json:"categoryId"`
	Thumbnail  string                `gorm:"column:thumbnail;type:VARCHAR(256);" json:"thumbnail"`
	IsTop      string                `gorm:"column:is_top;type:CHAR(1);" json:"isTop"`
	Status     string                `gorm:"column:status;type:CHAR(1);" json:"status"`
	ViewCount  int64                 `gorm:"column:view_count;type:BIGINT(200);" json:"viewCount"`
	IsComment  string                `gorm:"column:is_comment;type:CHAR(1);" json:"isComment"`
	CreateBy   int64                 `gorm:"column:create_by;type:BIGINT(20);" json:"createBy"`
	CreateTime time.Time             `gorm:"column:create_time;type:DATETIME;default:current_timestamp" json:"createTime"`
	UpdateBy   int64                 `gorm:"column:update_by;type:BIGINT(20);" json:"updateBy"`
	UpdateTime time.Time             `gorm:"column:update_time;type:DATETIME;default:current_timestamp on update current_timestamp" json:"updateTime"`
	DelFlag    soft_delete.DeletedAt `gorm:"column:del_flag;type:INT(1);softDelete:flag" json:"delFlag"` // 0：表示未删除的， 1：表示已经删除
}

func (SgArticle) TableName() string {
	return "sg_article"
}
