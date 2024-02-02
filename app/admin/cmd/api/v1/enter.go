package v1

import (
	"sgblog-go/app/admin/cmd/api/v1/admin"
)

type ApiGroup struct {
	BlogApiGroup admin.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
