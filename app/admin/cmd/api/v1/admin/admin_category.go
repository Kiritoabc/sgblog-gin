package admin

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/common/response"
	"strconv"
)

type AdminCategoryApi struct{}

func (s *AdminCategoryApi) ListAllCategory(ctx *gin.Context) {
	listAllCategory, err := blogCategoryService.ListAllCategory()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(listAllCategory, "获取成功", ctx)
}

func (s *AdminCategoryApi) Edit(ctx *gin.Context) {
	var category blog.SgCategory
	// get params
	err := ctx.ShouldBindJSON(category)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	// update
	err = global.SG_BLOG_DB.Model(&blog.SgCategory{}).
		Where("id = ?", category.Id).
		Updates(category).Error
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("修改成功", ctx)
}

func (s *AdminCategoryApi) Remove(ctx *gin.Context) {
	// get path param
	id := ctx.Query("id")

	// remove by id
	err := global.SG_BLOG_DB.Delete(&blog.SgCategory{}, id).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithMessage("删除成功", ctx)
}

func (s *AdminCategoryApi) Add(ctx *gin.Context) {
	// get json params
	var category blog.SgCategory
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// use gorm create to create a new record
	err = global.SG_BLOG_DB.Create(&category).Error
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithMessage("添加成功", ctx)
}

func (s *AdminCategoryApi) GetInfo(ctx *gin.Context) {
	// get path param id
	id := ctx.Query("id")

	// use gorm to select by id
	var category blog.SgCategory
	err := global.SG_BLOG_DB.Model(&blog.SgCategory{}).
		Where("id = ?", id).
		First(&category).Error

	// handle err
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// return success
	response.OkWithDetailed(category, "获取成功", ctx)
}

func (s *AdminCategoryApi) List(ctx *gin.Context) {
	// get params
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	name := ctx.Query("name")
	status := ctx.Query("status")

	// apply service
	list, total, err := blogCategoryService.SelectCategoryPage(blog.SgCategory{
		Name:   name,
		Status: status,
	}, pageNum, pageSize)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithDetailed(response.PageResult{
		Rows:  list,
		Total: total,
	}, "操作成功", ctx)
}
