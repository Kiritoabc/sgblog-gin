package blog

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"sgblog-go/app/blog/cmd/global"
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

func (s BlogLoginApi) Logout(ctx *gin.Context) {
	loginUser, exists := ctx.Get("loginUser")
	if !exists {
		response.FailWithMessage("用户未登录", ctx)
		return
	}
	// 删除redis中的token
	userId := loginUser.(*blog.UserLogin).User.Id
	key := fmt.Sprintf("bloglogin:%d", userId)
	err := global.SG_BLOG_REDIS.Del(context.Background(), key).Err()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("成功退出", ctx)
}
