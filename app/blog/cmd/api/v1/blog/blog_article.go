package blog

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/model/common/response"
	"strconv"
)

type BlogArticleApi struct{}

func (s *BlogArticleApi) HotArticleList(ctx *gin.Context) {
	hotArticleList, err := blogArticleService.HotArticleList()
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.OkWithData(hotArticleList, ctx)
}

func (s *BlogArticleApi) ArticleList(ctx *gin.Context) {
	// 获取分页参数	// 获取分页参数
	pageNum, _ := ctx.GetQuery("pageNum")
	pageNumInt, _ := strconv.ParseInt(pageNum, 10, 64)
	pageSize, _ := ctx.GetQuery("pageSize")
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)
	// 获取查询参数
	categoryId, _ := ctx.GetQuery("categoryId")
	if categoryId == "" {
		categoryId = "0"
	}
	categoryIdInt, _ := strconv.ParseInt(categoryId, 10, 64)
	// 调用service层方法获取文章列表
	articleList, total, err := blogArticleService.ArticleList(int(pageNumInt), int(pageSizeInt), categoryIdInt)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(response.PageResult{
		Rows:  articleList,
		Total: total,
	}, ctx)
}

func (s BlogArticleApi) GetArticleDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	idParam, _ := strconv.ParseInt(id, 10, 64)
	detail, err := blogArticleService.GetArticleDetail(idParam)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(detail, ctx)
}
