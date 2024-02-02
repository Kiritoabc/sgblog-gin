package blog

import (
	"context"
	"github.com/jinzhu/copier"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/vo"
	"sgblog-go/common/constants"
	"strconv"
)

type ArticleService struct{}

func (s *ArticleService) HotArticleList() ([]*vo.HotArticleVo, error) {
	// 1.查询热门文章
	// 2.必须是正式文章
	// 3.按照流量进行排序
	// 4.最多10条
	var articles []*blog.SgArticle
	db := global.SG_BLOG_DB
	err := db.Model(blog.SgArticle{}).
		Where("status = ?", constants.Normal).
		Order("view_count desc").
		Limit(10).
		Find(&articles).Error
	if err != nil {
		return nil, err
	}
	var vos []*vo.HotArticleVo
	err = copier.Copy(&vos, &articles)
	if err != nil {
		return nil, err
	}
	return vos, nil
}

func (s *ArticleService) ArticleList(pageNum int, pageSize int,
	categoryId int64) ([]*vo.ArticleVoList, int64, error) {
	// 1.查询，正式发布的，对isTop进行降序
	articleDB := global.SG_BLOG_DB.Model(blog.SgArticle{})

	if categoryId > 0 {
		articleDB = articleDB.Where("category_id = ?", categoryId)
	}
	var articles []*blog.SgArticle
	var count int64
	err := articleDB.Where("status = ?", constants.ArticleStatusNormal).
		Count(&count).
		Order("is_top desc").
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&articles).Error

	if err != nil {
		return nil, 0, err
	}
	var vos []*vo.ArticleVoList
	err = copier.Copy(&vos, articles)
	if err != nil {
		return nil, 0, err
	}
	return vos, count, nil
}

func (s *ArticleService) GetArticleDetail(id int64) (*vo.ArticleDetailVo, error) {
	// 1. 根据id查询文章
	db := global.SG_BLOG_DB.Model(blog.SgArticle{})
	var article *blog.SgArticle
	err := db.Where("id = ?", id).First(&article).Error
	if err != nil {
		return nil, err
	}

	// 2.从redis中获取viewCount
	redis := global.SG_BLOG_REDIS
	result := redis.HMGet(context.Background(), "article:viewCount", strconv.FormatInt(id, 10)).Val()
	str := string(result[0].(string))
	viewCount, _ := strconv.ParseInt(str, 10, 64)
	article.ViewCount = viewCount

	var articleDetailVo = &vo.ArticleDetailVo{}
	err = copier.Copy(&articleDetailVo, &article)
	if err != nil {
		return nil, err
	}

	// 3.根据分类id查询分类名
	var categoryId int64 = articleDetailVo.CategoryId
	var category *blog.SgCategory
	err = global.SG_BLOG_DB.Model(blog.SgCategory{}).
		Where("id = ?", categoryId).
		First(&category).Error
	if err != nil {
		return nil, err
	}
	if category != nil {
		articleDetailVo.CategoryName = category.Name
	}
	return articleDetailVo, nil
}

func (s *ArticleService) UpdateViewCount(id int64) error {
	redis := global.SG_BLOG_REDIS
	_, err := redis.
		HIncrBy(context.Background(), "article:viewCount", strconv.FormatInt(id, 10), 1).
		Result()
	if err != nil {
		return err
	}
	return nil
}
