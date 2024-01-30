package initialize

import (
	"github.com/robfig/cron/v3"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/blog/cmd/job"
)

func Timer() {
	if global.SG_BLOG_COFIG.Timer.Start {
		var option []cron.Option
		// 注册
		_, err := global.SG_BLOG_Timer.AddTaskByFunc("UpdateViewCount",
			global.SG_BLOG_COFIG.Timer.Spec, job.UpdateViewCountJob, option...)
		if err != nil {
			global.SG_BLOG_LOG.Error("定时任务注册失败")
		}
	}
}
