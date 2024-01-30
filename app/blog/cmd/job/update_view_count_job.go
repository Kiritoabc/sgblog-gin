package job

import (
	"context"
	"fmt"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/model/blog"
	"strconv"
)

func UpdateViewCountJob() {
	// 每55s 执行一次
	//time.Sleep(55 * time.Second)
	redis := global.SG_BLOG_REDIS
	global.SG_BLOG_LOG.Info("===========更新文章阅读量,更新文章阅读量开始=============")
	result, err := redis.HGetAll(context.Background(), "article:viewCount").Result()
	if err != nil {
		global.SG_BLOG_LOG.Error("更新文章阅读量失败")
	}
	// 更新文章阅读量
	var articles []*blog.SgArticle
	for k, v := range result {
		fmt.Printf("%s,%s", k, v)
		id, _ := strconv.ParseInt(k, 10, 64)
		viewCount, _ := strconv.ParseInt(v, 10, 64)
		var article = &blog.SgArticle{
			Id:        id,
			ViewCount: viewCount,
		}
		articles = append(articles, article)
	}
	// todo: 暂时没有看到批量更新
	for _, article := range articles {
		err = global.SG_BLOG_DB.Model(&blog.SgArticle{}).Where("id = ?", article.Id).Updates(article).Error
		if err != nil {
			global.SG_BLOG_LOG.Error("更新文章阅读量失败")
		}
	}

}
