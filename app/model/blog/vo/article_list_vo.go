package vo

import "time"

type ArticleVoList struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`        // 标题
	Summary      string    `json:"summary"`      // 文章摘要
	CategoryName string    `json:"categoryName"` // 所属分类
	Thumbnail    string    `json:"thumbnail"`    // 缩略图
	ViewCount    int64     `json:"viewCount"`    // 浏览量
	CreateTime   time.Time `json:"createTime"`
}
