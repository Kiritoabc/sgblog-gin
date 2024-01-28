package vo

type HotArticleVo struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	ViewCount int64  `json:"viewCount"`
}
