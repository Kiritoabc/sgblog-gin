package vo

type CategoryVo struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryVoList struct {
	List []*CategoryVo `json:"list"`
}
