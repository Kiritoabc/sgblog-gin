package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"sgblog-go/global"
)

func Redis() {
	redisCfg := global.SG_BLOG_COFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.SG_BLOG_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.SG_BLOG_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GVA_REDIS = client
	}
}
