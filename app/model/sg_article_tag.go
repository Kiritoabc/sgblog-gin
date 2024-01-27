package model

type SgArticleTag struct {
	ArticleId int64 `gorm:"column:article_id;type:BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	TagId     int64 `gorm:"column:tag_id;type:BIGINT(20);NOT NULL"`
}
