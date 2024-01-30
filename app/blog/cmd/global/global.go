package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"sgblog-go/common/config"
	"sgblog-go/common/utils"
	"sync"
)

var (
	SG_BLOG_DB                   *gorm.DB
	SG_BLOG_DBList               map[string]*gorm.DB
	SG_BLOG_COFIG                config.Server
	SG_BLOG_VP                   *viper.Viper
	SG_BLOG_LOG                  *zap.Logger
	lock                         sync.RWMutex
	SG_BLOG_REDIS                *redis.Client
	SG_BLOG__Concurrency_Control = &singleflight.Group{}
	SG_BLOG_Timer                = utils.NewTimerTask()
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return SG_BLOG_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := SG_BLOG_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
