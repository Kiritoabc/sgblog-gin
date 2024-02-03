package admin

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/dto"
	"sgblog-go/app/model/common/response"
	"strconv"
)

type AdminArticleApi struct {
}

func (s *AdminArticleApi) Add(ctx *gin.Context) {
	var articleDto dto.AddArticleDto
	err := ctx.ShouldBind(&articleDto)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = blogArticleService.Add(articleDto)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("添加成功", ctx)
}

func (s *AdminArticleApi) Edit(ctx *gin.Context) {
	var articleDto dto.AddArticleDto
	err := ctx.ShouldBind(&articleDto)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = blogArticleService.Edit(articleDto)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("修改成功", ctx)
}

func (s *AdminArticleApi) Delete(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err = blogArticleService.Delete(id); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("删除成功", ctx)
}

func (s *AdminArticleApi) GetInfo(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	articleVo, _ := blogArticleService.GetInfo(id)
	response.OkWithDetailed(articleVo, "获取成功", ctx)
}

func (s *AdminArticleApi) List(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	title := ctx.Query("title")
	summary := ctx.Query("summary")

	list, total, err := blogArticleService.List(blog.SgArticle{Title: title, Summary: summary},
		pageNum, pageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(response.PageResult{
		Rows:  list,
		Total: total,
	}, "获取成功", ctx)
}
