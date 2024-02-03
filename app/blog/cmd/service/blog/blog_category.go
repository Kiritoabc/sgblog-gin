package blog

import (
	"github.com/jinzhu/copier"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/vo"
	"sgblog-go/common/constants"
	"sync"
)

type CategoryService struct{}

func (s *CategoryService) GetCategoryList() (*vo.CategoryVoList, error) {
	// 1.查询文章表 状态为已发布的文章
	db := global.SG_BLOG_DB
	var articleList []*blog.SgArticle
	err := db.Model(&blog.SgArticle{}).Where("status = ?", constants.ArticleStatusNormal).Find(&articleList).Error
	if err != nil {
		return nil, err
	}
	// 2.获取文章的分类id，并且去重
	categoryIds := getCategoryIdsSet(articleList)
	// 3.查询分类表
	var categories []*blog.SgCategory
	err = db.Model(&blog.SgCategory{}).
		Where("id in (?) and status = ?", categoryIds, constants.STATUS_NORMAL).
		Find(&categories).Error
	if err != nil {
		return nil, err
	}
	// 4.封装vo
	var cateroryVo = vo.CategoryVoList{}
	err = copier.Copy(&cateroryVo.List, &categories)
	if err != nil {
		return nil, err
	}
	return &cateroryVo, nil
}

func getCategoryIdsSet(articleList []*blog.SgArticle) []int64 {
	categoryIdsMap := extractCategoryIds(articleList)
	categoryIds := make([]int64, 0, len(categoryIdsMap))
	for id := range categoryIdsMap {
		categoryIds = append(categoryIds, id)
	}
	return categoryIds
}

func extractCategoryIds(articleList []*blog.SgArticle) map[int64]struct{} {
	categoryIds := make(map[int64]struct{})
	var mutex sync.Mutex // 如果需要并发安全，可以添加互斥锁
	for _, article := range articleList {
		mutex.Lock()
		categoryIds[article.CategoryId] = struct{}{}
		mutex.Unlock()
	}

	return categoryIds
}
