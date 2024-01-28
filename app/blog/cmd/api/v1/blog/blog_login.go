package blog

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/common/response"
)

type BlogLoginApi struct{}

func (s *BlogLoginApi) Login(ctx *gin.Context) {
	var user = blog.SysUser{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	blogUserLoginVo, err := blogLoginService.Login(user)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(blogUserLoginVo, ctx)
}
