package admin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/vo"
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

func (s *AdminLoginApi) Logout(ctx *gin.Context) {
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
	loginUser, exists := ctx.Get("loginUser")
	if !exists {
		response.FailWithMessage("用户未登录", ctx)
		return
	}
	var isAdmin = false
	if loginUser.(*blog.UserLogin).User.Type == "1" {
		isAdmin = true
	}
	userId := loginUser.(*blog.UserLogin).User.Id
	//根据用户id查询权限信息
	perms, err := menuService.SelectPermsByUserId(userId, isAdmin)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	//根据用户id查询角色信息
	roleKeyList, _ := roleService.SelectRoleKeyByUserId(userId)

	var userInfoVo = vo.UserInfoVo{}
	err = copier.Copy(&userInfoVo, &loginUser.(*blog.UserLogin).User)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	adminUserInfoVo := vo.AdminUserInfoVo{
		Permissions: perms,
		Roles:       roleKeyList,
		UserInfoVo:  userInfoVo,
	}
	response.OkWithDetailed(adminUserInfoVo, "获取成功", ctx)
}

func (s *AdminLoginApi) GetRouters(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		response.FailWithMessage("用户未登录", ctx)
		return
	}
	// 查询menu 结果是tree的形式
	_ = userId
}
