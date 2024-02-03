package dto

type AddArticleDto struct {
	ID         int64   `json:"id"`
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	Summary    string  `json:"summary"`
	CategoryID int64   `json:"categoryId"`
	Thumbnail  string  `json:"thumbnail"`
	IsTop      string  `json:"isTop"` // 注意：这里假设isTop、status和isComment在数据库或API中存储的是字符串类型
	Status     string  `json:"status"`
	ViewCount  int64   `json:"viewCount"`
	IsComment  string  `json:"isComment"`
	Tags       []int64 `json:"tags"`
}
