package blog

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/model/common/response"
)

type BlogArticleApi struct{}

func (s BlogArticleApi) HotArticleList(ctx *gin.Context) {
	hotArticleList, err := blogArticleService.HotArticleList()
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.OkWithData(hotArticleList, ctx)
}
