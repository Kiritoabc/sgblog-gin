package vo

type RoleMenuTreeSelectVo struct {
	CheckedKeys []int64 `json:"checkedKeys"`

	Menus []*MenuTreeVo `json:"menus"`
}
