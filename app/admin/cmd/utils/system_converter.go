package utils

import (
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/vo"
)

func NewMenuTreeVo(id int64, name string, parentId int64) *vo.MenuTreeVo {
	return &vo.MenuTreeVo{
		Id:       id,
		Label:    name,
		ParentId: parentId,
		Children: []*vo.MenuTreeVo{},
	}
}

func SystemConverter(menus []*blog.SysMenu) []*vo.MenuTreeVo {
	menuTreeVos := make([]*vo.MenuTreeVo, 0)
	for _, m := range menus {
		menuTreeVos = append(menuTreeVos, NewMenuTreeVo(m.Id, m.MenuName, m.ParentId))
	}

	options := make([]*vo.MenuTreeVo, 0)
	for _, o := range menuTreeVos {
		if o.ParentId == 0 {
			o.Children = getChildList(menuTreeVos, o)
			options = append(options, o)
		}
	}

	return options
}

func getChildList(list []*vo.MenuTreeVo, option *vo.MenuTreeVo) []*vo.MenuTreeVo {
	var children []*vo.MenuTreeVo
	for _, o := range list {
		if o.ParentId == option.Id {
			o.Children = getChildList(list, o)
			children = append(children, o)
		}
	}
	return children
}
