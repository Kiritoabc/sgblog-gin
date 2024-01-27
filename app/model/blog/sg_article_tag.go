package blog

type SgArticleTag struct {
	ArticleId int64 `gorm:"column:article_id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL" json:"articleId"`
	TagId     int64 `gorm:"column:tag_id;type:BIGINT(20);NOT NULL" json:"tagId"`
}

func (SgArticleTag) TableName() string {
	return "sg_article_tag"
}
