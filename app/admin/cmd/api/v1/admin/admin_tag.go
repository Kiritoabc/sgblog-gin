package admin

import (
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/dto"
	"sgblog-go/app/model/common/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminTagApi struct {
}

func (s *AdminTagApi) List(ctx *gin.Context) {
	// get params
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	name := ctx.Query("name")
	remark := ctx.Query("remark")

	list, tatal, err := blogTagService.PageTagList(dto.TagListDto{
		Name:   name,
		Remark: remark,
	}, pageNum, pageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithDetailed(response.PageResult{
		Rows:  list,
		Total: tatal,
	}, "操作成功", ctx)
}

func (s *AdminTagApi) Add(ctx *gin.Context) {
	var tag blog.SgTag

	// get json param
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = global.SG_BLOG_DB.Create(&tag).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithMessage("添加成功", ctx)
}

func (s *AdminTagApi) Delete(ctx *gin.Context) {
	id := ctx.Query("id")

	// remove by id
	if err := global.SG_BLOG_DB.Delete(&blog.SgTag{}, id).Error; err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// return success
	response.OkWithMessage("删除成功", ctx)
}

func (s *AdminTagApi) Edit(ctx *gin.Context) {
	var tag blog.SgTag
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = global.SG_BLOG_DB.Model(&blog.SgTag{}).
		Where("id = ?", tag.Id).
		Updates(&tag).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("修改成功", ctx)
}

func (s *AdminTagApi) GetInfo(ctx *gin.Context) {
	// get id
	id := ctx.Query("id")

	// select by id
	var tag blog.SgTag
	err := global.SG_BLOG_DB.Model(&blog.SgTag{}).
		Where("id = ?", id).
		First(&tag).Error

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(tag, "查询成功", ctx)
}

func (s *AdminTagApi) ListAllTag(ctx *gin.Context) {
	tags, err := blogTagService.ListAllTag()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(tags, "查询成功", ctx)
}
