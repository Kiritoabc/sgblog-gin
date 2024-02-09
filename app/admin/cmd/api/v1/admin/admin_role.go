package admin

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/dto"
	"sgblog-go/app/model/common/response"
	"strconv"
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

func (s *AdminRoleApi) GetInfo(ctx *gin.Context) {
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

func (s *AdminRoleApi) Add(ctx *gin.Context) {
	var role blog.SysRole

	err := ctx.ShouldBindJSON(&role)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = roleService.InsertRole(role)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
}

func (s *AdminRoleApi) List(ctx *gin.Context) {
	roleName := ctx.Query("role_name")
	status := ctx.Query("status")

	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	list, total, err := roleService.List(blog.SysRole{RoleName: roleName, Status: status}, pageNum, pageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(response.PageResult{Rows: list, Total: total}, "查询成功", ctx)
}

func (s *AdminRoleApi) ChangeStatus(ctx *gin.Context) {
	var changeRoleStatusDto dto.ChangeRoleStatusDto

	err := ctx.ShouldBind(&changeRoleStatusDto)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// change status
	err = global.SG_BLOG_DB.
		Model(&blog.SysRole{}).
		Where("id = ?", changeRoleStatusDto.RoleId).
		Updates(blog.SysRole{Status: changeRoleStatusDto.Status}).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("修改成功", ctx)
}
