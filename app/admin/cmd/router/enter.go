package router

import "sgblog-go/app/admin/cmd/router/admin"

type RouterGroup struct {
	Admin admin.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
