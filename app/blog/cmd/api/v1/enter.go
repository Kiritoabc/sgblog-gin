package v1

import "sgblog-go/app/blog/cmd/api/v1/blog"

type ApiGroup struct {
	BlogApiGroup blog.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
