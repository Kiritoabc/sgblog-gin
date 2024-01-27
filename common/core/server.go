package core

import (
	"fmt"
	"go.uber.org/zap"
	"sgblog-go/common/global"
	initialize2 "sgblog-go/common/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.SG_BLOG_COFIG.System.UseMultipoint || global.SG_BLOG_COFIG.System.UseRedis {
		// 初始化redis服务
		//todo: 暂时没有配置redis
		initialize2.Redis()
	}

	// 从db加载jwt数据
	//if global.GVA_DB != nil {
	//	system.LoadAll()
	//}

	Router := initialize2.Routers()
	//Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.SG_BLOG_COFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.SG_BLOG_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.SG_BLOG_LOG.Error(s.ListenAndServe().Error())
}
