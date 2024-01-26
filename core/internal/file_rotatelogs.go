package internal

import (
	"os"
	"sgblog-go/global"

	"go.uber.org/zap/zapcore"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *fileRotatelogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := NewCutter(global.SG_BLOG_COFIG.Zap.Director, level, WithCutterFormat("2006-01-02"))
	if global.SG_BLOG_COFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
