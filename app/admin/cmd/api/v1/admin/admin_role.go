package admin

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/common/response"
)

type AdminRoleApi struct {
}

func (s *AdminRoleApi) ListAllRole(ctx *gin.Context) {
	roleAll, err := roleService.SelectRoleAll()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(roleAll, "获取角色成功", ctx)
}

func (s *AdminRoleApi) getInfo(ctx *gin.Context) {
	var role blog.SysRole
	roleId := ctx.Param("roleId")

	err := global.SG_BLOG_DB.Model(&blog.SysRole{}).Where("role_id = ?", roleId).First(&role).Error
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(role, "获取角色成功", ctx)
}

func (s *AdminRoleApi) Edit(ctx *gin.Context) {

	var role blog.SysRole
	err := ctx.ShouldBindJSON(&role)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = roleService.UpdateRole(role)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithMessage("修改成功", ctx)
}

func (s *AdminRoleApi) Remove(ctx *gin.Context) {
	id := ctx.Param("id")

	err := global.SG_BLOG_DB.Where("id = ?", id).Delete(&blog.SysRole{}).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithMessage("删除成功", ctx)
}

// TODO

func (s *AdminRoleApi) Add(ctx *gin.Context) {
	var role blog.SysRole

	err := ctx.ShouldBindJSON(&role)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
}

func (s *AdminRoleApi) List(ctx *gin.Context) {

}

func (s *AdminRoleApi) ChangeStatus(ctx *gin.Context) {

}
