package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/dto"
	"sgblog-go/app/model/common/response"
	"sgblog-go/common/constants"
	"strconv"
)

type BlogCommentApi struct{}

func (s BlogCommentApi) CommentList(ctx *gin.Context) {
	// 获取参数并转换
	articleIdStr := ctx.Query("articleId")
	pageNumStr := ctx.Query("pageNum")
	pageSizeStr := ctx.Query("pageSize")
	var articleId int64
	var pageNum, pageSize int

	articleId, _ = strconv.ParseInt(articleIdStr, 10, 64)
	pageNum, _ = strconv.Atoi(pageNumStr)
	pageSize, _ = strconv.Atoi(pageSizeStr)

	commentList, total, err := blogCommentService.CommentList(constants.ArticleComment, articleId, pageNum, pageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(response.PageResult{
		Rows:  commentList,
		Total: total,
	}, "操作成功", ctx)
}

func (s *BlogCommentApi) AddComment(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		response.FailWithMessage("用户未登录", ctx)
		return
	}
	var addCommentDto dto.AddCommentDto
	err := ctx.ShouldBind(&addCommentDto)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	var comment blog.SgComment
	err = copier.Copy(&comment, &addCommentDto)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	comment.CreateBy = userId.(int64)
	comment.UpdateBy = userId.(int64)
	err = blogCommentService.AddComment(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("操作成功", ctx)
}

func (s *BlogCommentApi) LinkCommentList(ctx *gin.Context) {
	pageNumStr := ctx.Query("pageNum")
	pageSizeStr := ctx.Query("pageSize")

	var pageNum, pageSize int
	pageNum, _ = strconv.Atoi(pageNumStr)
	pageSize, _ = strconv.Atoi(pageSizeStr)

	linkCommentList, total, err := blogCommentService.CommentList(constants.LinkComment, -1, pageNum, pageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithDetailed(response.PageResult{
		Rows:  linkCommentList,
		Total: total,
	}, "操作成功", ctx)
}
