package blog

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/model/common/response"
)

type BlogCategoryApi struct{}

func (s *BlogCategoryApi) GetCategoryList(ctx *gin.Context) {

	categoryVo, err := blogCategoryService.GetCategoryList()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(categoryVo.List, ctx)
}
