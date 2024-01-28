package vo

type ArticleDetailVo struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`        // 标题
	Summary      string `json:"summary"`      // 摘要
	CategoryId   int64  `json:"categoryId"`   // 所属分类id
	CategoryName string `json:"categoryName"` // 所属分类名
	Thumbnail    string `json:"thumbnail"`    //缩略图
	Content      string `json:"content"`      // 文章内容
	ViewCount    int64  `json:"viewCount"`    // 浏览量
	CreateTime   string `json:"createTime"`   // 创建时间
}
