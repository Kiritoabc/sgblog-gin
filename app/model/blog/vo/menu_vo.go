package vo

type MenuVo struct {
	ID        int64  `json:"id"`
	MenuName  string `json:"menuName"`
	ParentID  int64  `json:"parentId"`
	OrderNum  int    `json:"orderNum"`
	Path      string `json:"path"`
	Component string `json:"component"`
	IsFrame   int    `json:"isFrame,string"` // 使用string防止数字0被解析成false
	MenuType  string `json:"menuType"`
	Visible   string `json:"visible"`
	Status    string `json:"status"`
	Perms     string `json:"perms"`
	Icon      string `json:"icon"`
	Remark    string `json:"remark"`
}
