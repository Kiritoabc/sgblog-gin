package admin

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/dto"
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
	// 查询数量
	articleDB.Where("status = ?", constants.ArticleStatusNormal).Count(&count)

	err := articleDB.Where("status = ?", constants.ArticleStatusNormal).
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

func (s *ArticleService) Add(articleDto dto.AddArticleDto) error {
	var article blog.SgArticle
	err := copier.Copy(&article, &articleDto)
	if err != nil {
		return err
	}
	// 开启事务
	db := global.SG_BLOG_DB.Begin()
	if err = db.Create(article).Error; err != nil {
		if err := db.Rollback().Error; err != nil {
			global.SG_BLOG_LOG.Error("Add article rollback error", zap.Error(err))
		}
		return err
	}
	// 创建ArticleTag关联记录
	articleTags := make([]blog.SgArticleTag, 0, len(articleDto.Tags))
	for _, tagId := range articleDto.Tags {
		articleTags = append(articleTags, blog.SgArticleTag{ArticleId: article.Id, TagId: tagId})
	}

	// 批量保存关联数据
	if err = db.Create(&articleTags).Error; err != nil {
		if err := db.Rollback().Error; err != nil {
			global.SG_BLOG_LOG.Error("Save articleTags rollback error", zap.Error(err))
		}
		return err
	}
	// 事务提交
	if err = db.Commit().Error; err != nil {
		fmt.Println("Failed to commit transaction:", err)
	} else {
		fmt.Println("Transaction committed successfully.")
	}
	return nil
}

func (s *ArticleService) Edit(articleDto dto.AddArticleDto) error {
	var article blog.SgArticle
	err := copier.Copy(&article, &articleDto)
	if err != nil {
		return err
	}
	// 开启事务
	db := global.SG_BLOG_DB.Begin()
	if err = db.Save(article).Error; err != nil {
		if err := db.Rollback().Error; err != nil {
			global.SG_BLOG_LOG.Error("Update article rollback error", zap.Error(err))
		}
		return err
	}
	if err = db.Delete(&blog.SgArticleTag{}, articleDto.ID).Error; err != nil {
		if err := db.Rollback().Error; err != nil {
			global.SG_BLOG_LOG.Error("Delete articleTag rollback error", zap.Error(err))
		}
		return err
	}

	// 创建ArticleTag关联记录
	articleTags := make([]blog.SgArticleTag, 0, len(articleDto.Tags))
	for _, tagId := range articleDto.Tags {
		articleTags = append(articleTags, blog.SgArticleTag{ArticleId: article.Id, TagId: tagId})
	}

	// 批量保存关联数据
	if err = db.Create(&articleTags).Error; err != nil {
		if err := db.Rollback().Error; err != nil {
			global.SG_BLOG_LOG.Error("Save articleTags rollback error", zap.Error(err))
		}
		return err
	}
	// 事务提交
	if err = db.Commit().Error; err != nil {
		fmt.Println("Failed to commit transaction:", err)
	} else {
		fmt.Println("Transaction committed successfully.")
	}
	return nil
}

func (s *ArticleService) Delete(id int64) error {
	return global.SG_BLOG_DB.Delete(&blog.SgArticle{}, id).Error
}

func (s *ArticleService) GetInfo(id int64) (*vo.ArticleVo, error) {
	var article blog.SgArticle
	db := global.SG_BLOG_DB
	if err := db.Model(&blog.SgArticle{}).
		Where("id = ?", id).
		First(&article).Error; err != nil {
		return nil, err
	}
	// 获取关联标签
	var tags []blog.SgArticleTag
	if err := db.Model(&blog.SgArticleTag{}).
		Where("article_id = ?", id).
		Find(&tags).Error; err != nil {
		return nil, err
	}
	var tagsVo []int64
	for _, tag := range tags {
		tagsVo = append(tagsVo, tag.TagId)
	}
	var articleVo vo.ArticleVo
	err := copier.Copy(&articleVo, &article)
	if err != nil {
		return nil, err
	}
	articleVo.Tags = tagsVo
	return &articleVo, nil
}

func (s *ArticleService) List(article blog.SgArticle, pageNum int, pageSize int) ([]*blog.SgArticle, int64, error) {
	var articles []*blog.SgArticle

	query := global.SG_BLOG_DB.Model(&blog.SgArticle{}) // 初始化你的GORM数据库连接...

	// 构建查询条件
	if article.Title != "" {
		query = query.Where("title LIKE ?", "%"+article.Title+"%")
	}
	if article.Summary != "" {
		query = query.Where("summary LIKE ?", "%"+article.Summary+"%")
	}

	// 分页设置
	offset := (pageNum - 1) * pageSize
	limit := pageSize

	var total int64
	// 执行查询并获取结果
	if err := query.Count(&total).Offset(offset).Limit(limit).Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}
