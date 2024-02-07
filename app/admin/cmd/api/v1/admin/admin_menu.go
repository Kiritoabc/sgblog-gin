package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/admin/cmd/utils"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/vo"
	"sgblog-go/app/model/common/response"
	"strconv"
)

type AdminMenuApi struct{}

func (s *AdminMenuApi) TreeSelect(ctx *gin.Context) {

	list, err := menuService.SelectMenuList(blog.SysMenu{})

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	menuTreeVos := utils.SystemConverter(list)

	response.OkWithDetailed(menuTreeVos, "获取成功", ctx)
}

func (s *AdminMenuApi) RoleMenuTreeSelect(ctx *gin.Context) {
	roleId := ctx.Param("roleId")

	menus, err := menuService.SelectMenuList(blog.SysMenu{})
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	parseInt, _ := strconv.ParseInt(roleId, 10, 64)
	listByRoleId, _ := menuService.SelectMenuListByRoleId(parseInt)

	menuTreeVos := utils.SystemConverter(menus)

	vo := &vo.RoleMenuTreeSelectVo{
		CheckedKeys: listByRoleId,
		Menus:       menuTreeVos,
	}

	response.OkWithDetailed(vo, "获取成功", ctx)
}

func (s *AdminMenuApi) List(ctx *gin.Context) {
	menuName := ctx.Query("menu_name")
	status := ctx.Query("status")
	menuList, err := menuService.SelectMenuList(blog.SysMenu{
		MenuName: menuName,
		Status:   status,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	var menuVoList []*vo.MenuVo

	err = copier.Copy(&menuVoList, &menuList)
	// handle error
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithDetailed(menuVoList, "获取成功", ctx)
}

func (s *AdminMenuApi) Add(ctx *gin.Context) {
	var menu blog.SysMenu

	err := ctx.ShouldBindJSON(&menu)

	err = global.SG_BLOG_DB.Create(&menu).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithMessage("添加成功", ctx)
}

func (s *AdminMenuApi) GetInfo(ctx *gin.Context) {
	menuId := ctx.Param("menuId")
	var menu blog.SysMenu

	err := global.SG_BLOG_DB.Model(&blog.SysMenu{}).Where("menu_id = ?", menuId).First(&menu).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(&menu, "获取成功", ctx)
}

func (s *AdminMenuApi) Edit(ctx *gin.Context) {
	var menu blog.SysMenu
	err := ctx.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	if menu.Id == menu.ParentId {
		response.FailWithMessage("上级不能是自己", ctx)
		return
	}

	err = global.SG_BLOG_DB.Model(&blog.SysMenu{}).Where("menu_id = ?", menu.Id).Updates(&menu).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithMessage("修改成功", ctx)
}

func (s *AdminMenuApi) Remove(ctx *gin.Context) {
	menuId := ctx.Param("menuId")
	// 存在子菜单，不允许删除

	if menuService.HasChild(menuId) {
		response.FailWithMessage("存在子菜单，不允许删除", ctx)
		return
	}

	err := global.SG_BLOG_DB.Where("menu_id = ?", menuId).Delete(&blog.SysMenu{}).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithMessage("删除成功", ctx)
}
