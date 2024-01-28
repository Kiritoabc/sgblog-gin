package response

type PageResult struct {
	Rows  interface{} `json:"rows"`
	Total int64       `json:"total"`
}
