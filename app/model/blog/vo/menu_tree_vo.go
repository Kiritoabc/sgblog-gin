package vo

type MenuTreeVo struct {
	Id int64 `json:"id"`

	Label string `json:"label"`

	ParentId int64 `json:"parentId"`

	Children []*MenuTreeVo `json:"children"`
}
