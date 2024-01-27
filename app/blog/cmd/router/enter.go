package router

import "sgblog-go/app/blog/cmd/router/blog"

type RouterGroup struct {
	Blog blog.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
