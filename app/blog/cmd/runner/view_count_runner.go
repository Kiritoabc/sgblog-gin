package runner

import (
	"context"
	"go.uber.org/zap"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/model/blog"
)

func ViewCountRunner() {
	var articleList []*blog.SgArticle
	err := global.SG_BLOG_DB.Model(&blog.SgArticle{}).Find(&articleList).Error
	if err != nil {
		global.SG_BLOG_LOG.Error("获取文章列表失败", zap.Error(err))
	}
	for _, article := range articleList {
		err := global.SG_BLOG_REDIS.HSet(context.Background(), "article:viewCount", article.Id, article.ViewCount).Err()
		if err != nil {
			global.SG_BLOG_LOG.Error("更新文章阅读量失败", zap.Error(err))
		}
	}
}
