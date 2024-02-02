package admin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/common/response"
)

type AdminLoginApi struct{}

func (s *AdminLoginApi) Login(ctx *gin.Context) {
	var user = blog.SysUser{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	token, err := blogLoginService.Login(user)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp := map[string]string{"token": token}
	response.OkWithDetailed(resp, "登录成功", ctx)
}

func (s AdminLoginApi) Logout(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		response.FailWithMessage("用户未登录", ctx)
		return
	}
	// 删除redis中的token
	key := fmt.Sprintf("login:%d", userId)
	err := global.SG_BLOG_REDIS.Del(context.Background(), key).Err()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("退出登录成功", ctx)
}

func (s *AdminLoginApi) GetInfo(ctx *gin.Context) {
	ctx.Get("loginUser")
}
