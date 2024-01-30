package blog

import (
	"github.com/gin-gonic/gin"
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
