package admin

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/common/response"
	"strconv"
)

type AdminLinkApi struct{}

func (s *AdminLinkApi) List(ctx *gin.Context) {
	// get params
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	name := ctx.Query("name")
	status := ctx.Query("status")

	// apply service
	links, total, err := blogLinkService.SelectLinkPage(blog.SgLink{
		Name:   name,
		Status: status,
	}, pageNum, pageSize)

	// handle err
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// return success
	response.OkWithDetailed(response.PageResult{
		Rows:  links,
		Total: total,
	}, "查询成功", ctx)
}

func (s *AdminLinkApi) Add(ctx *gin.Context) {
	var link blog.SgLink

	// get json param
	err := ctx.ShouldBindJSON(&link)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = global.SG_BLOG_DB.Create(&link).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithMessage("添加成功", ctx)
}

func (s *AdminLinkApi) Delete(ctx *gin.Context) {
	id := ctx.Query("id")

	// remove by id
	if err := global.SG_BLOG_DB.Delete(&blog.SgLink{}, id).Error; err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// return success
	response.OkWithMessage("删除成功", ctx)
}

func (s *AdminLinkApi) Edit(ctx *gin.Context) {
	var link blog.SgLink
	err := ctx.ShouldBindJSON(&link)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = global.SG_BLOG_DB.Model(&blog.SgLink{}).
		Where("id = ?", link.Id).
		Updates(&link).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("修改成功", ctx)
}

func (s *AdminLinkApi) ChangeLinkStatus(ctx *gin.Context) {
	var link blog.SgLink
	err := ctx.ShouldBindJSON(&link)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = global.SG_BLOG_DB.Model(&blog.SgLink{}).
		Where("id = ?", link.Id).
		Updates(&link).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("修改成功", ctx)
}

func (s *AdminLinkApi) GetInfo(ctx *gin.Context) {
	// get id
	id := ctx.Query("id")

	// select by id
	var link blog.SgLink
	err := global.SG_BLOG_DB.Model(&blog.SgLink{}).
		Where("id = ?", id).
		First(&link).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(link, "查询成功", ctx)
}
