package blog

import (
	"github.com/gin-gonic/gin"
	"sgblog-go/app/model/common/response"
)

type BlogCategoryApi struct{}

func (s *BlogCategoryApi) GetCategoryList(ctx *gin.Context) {

	categoryVo, _ := blogCategoryService.GetCategoryList()
	response.OkWithData(categoryVo.List, ctx)
}
