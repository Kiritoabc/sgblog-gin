package service

import (
	"sgblog-go/app/service/blog"
)

type ServiceGroup struct {
	BlogServiceGroup blog.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
