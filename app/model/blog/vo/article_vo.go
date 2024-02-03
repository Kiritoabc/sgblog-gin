package vo

type ArticleVo struct {
	ID         int64   `json:"id"`
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	Summary    string  `json:"summary"`
	CategoryID int64   `json:"categoryId"`
	Thumbnail  string  `json:"thumbnail"`
	IsTop      string  `json:"isTop"` // 假设isTop、status和isComment是字符串类型
	Status     string  `json:"status"`
	ViewCount  int64   `json:"viewCount"`
	IsComment  string  `json:"isComment"`
	Tags       []int64 `json:"tags"`
}
