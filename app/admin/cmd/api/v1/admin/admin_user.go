package admin

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/vo"
	"sgblog-go/app/model/common/response"
	"strconv"
	"strings"
)

type AdminUserApi struct {
}

func (s *AdminUserApi) List(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	userName := ctx.Query("userName")
	status := ctx.Query("status")
	phonenumber := ctx.Query("phonenumber")

	list, total, err := blogUserService.SelectUserPage(blog.SysUser{
		UserName:    userName,
		Status:      status,
		Phonenumber: phonenumber},
		pageNum, pageSize)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(response.PageResult{
		Rows:  list,
		Total: total,
	}, "查询成功", ctx)
}

func (s *AdminUserApi) Add(ctx *gin.Context) {
	var user blog.SysUser
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if !blogUserService.CheckUserNameUnique(user) {
		response.FailWithMessage("添加失败，用户名已存在", ctx)
		return
	}
	if !blogUserService.CheckPhoneUnique(user) {
		response.FailWithMessage("添加失败，手机号已存在", ctx)
		return
	}
	if !blogUserService.CheckEmailUnique(user) {
		response.FailWithMessage("添加失败，邮箱已存在", ctx)
		return
	}
	err = global.SG_BLOG_DB.Create(&user).Error
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("添加成功", ctx)
}

/**
 * 根据用户编号获取详细信息
 */

func (s *AdminUserApi) GetUserInfoAndRoleIds(ctx *gin.Context) {
	userId := ctx.Param("userId")

	id, _ := strconv.ParseInt(userId, 10, 64)

	roles, err := roleService.SelectRoleAll()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	user, _ := blogUserService.GetById(id)

	roleIds, _ := roleService.SelectRoleIdByUserId(id)

	response.OkWithDetailed(vo.UserInfoAndRoleIdsVo{
		User:    user,
		Roles:   roles,
		RoleIds: roleIds,
	}, "获取用户详细信息成功", ctx)
}

func (s *AdminUserApi) Edit(ctx *gin.Context) {
	var user blog.SysUser

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = blogUserService.UpdateUser(user)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithMessage("修改成功", ctx)
}

/**
 * 删除用户
 */

func (s *AdminUserApi) Remove(ctx *gin.Context) {
	ids := ctx.Param("userIds")
	spiltStr := strings.Split(ids, ",")
	userIdsSlice := make([]int, len(spiltStr))

	for i, s := range spiltStr {
		id, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		userIdsSlice[i] = id
	}
	_ = ids
}
