package service

import (
	"sgblog-go/app/admin/cmd/service/admin"
)

type ServiceGroup struct {
	BlogServiceGroup admin.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
