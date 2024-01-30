package blog

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/model/common/response"
)

type BlogLinkApi struct{}

func (s *BlogLinkApi) GetAllLink(ctx *gin.Context) {
	allLink, err := blogLinkService.GetAllLink()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(allLink, "获取成功", ctx)
}
