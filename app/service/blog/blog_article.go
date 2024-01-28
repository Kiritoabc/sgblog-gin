package blog

import (
	"github.com/jinzhu/copier"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/vo"
	"sgblog-go/common/constants"
)

type ArticleService struct{}

func (service *ArticleService) HotArticleList() ([]*vo.HotArticleVo, error) {
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
