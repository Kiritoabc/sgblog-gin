package blog

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/common/response"
)

type BlogUserApi struct{}

func (s BlogUserApi) UserInfo(ctx *gin.Context) {
	loginUser, exists := ctx.Get("loginUser")
	if !exists {
		response.FailWithMessage("用户未登录", ctx)
		return
	}
	userId := loginUser.(*blog.UserLogin).User.Id
	userInfo, err := blogUserService.UserInfo(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithDetailed(userInfo, "操作成功", ctx)
}

func (s BlogUserApi) UpdateUserInfo(ctx *gin.Context) {
	var user blog.SysUser
	err := ctx.ShouldBind(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = blogUserService.UpdateUserInfo(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("操作成功", ctx)
}

func (s BlogUserApi) Register(ctx *gin.Context) {
	var user blog.SysUser
	err := ctx.ShouldBind(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = blogUserService.Register(user)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("操作成功", ctx)
}
