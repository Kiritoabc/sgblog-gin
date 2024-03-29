package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/blog/cmd/initialize/internal"
	"sgblog-go/common/config"
	"time"
)

// GormMysql 初始化Mysql数据库

func GormMysql() *gorm.DB {
	m := global.SG_BLOG_COFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		// 设置全局钩子
		db.Callback().Create().Before("gorm:create").Register("setup_create_time", func(tx *gorm.DB) {
			for _, field := range tx.Statement.Schema.Fields {
				if field.DBName == "create_time" {
					tx.Statement.SetColumn(field.DBName, time.Now())
				}
				if field.DBName == "update_time" {
					tx.Statement.SetColumn(field.DBName, time.Now())
				}
			}
		})
		//
		db.Callback().Update().Before("gorm:update").Register("setup_update_time", func(tx *gorm.DB) {
			for _, field := range tx.Statement.Schema.Fields {
				if field.DBName == "update_time" {
					tx.Statement.SetColumn(field.DBName, time.Now())
					break
				}
			}
		})

		return db
	}
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		// 设置数据库连接池
		sqlDB.SetConnMaxIdleTime(time.Hour)
		sqlDB.SetConnMaxLifetime(24 * time.Hour)
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
