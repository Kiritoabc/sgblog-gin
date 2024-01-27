package service

import "sgblog-go/app/blog/cmd/service/blog"

type ServiceGroup struct {
	BlogServiceGroup blog.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
