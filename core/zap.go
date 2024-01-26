package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sgblog-go/core/internal"
	"sgblog-go/global"
	"sgblog-go/utils"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.SG_BLOG_COFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory \n", global.SG_BLOG_COFIG.Zap.Director)
		err := os.Mkdir(global.SG_BLOG_COFIG.Zap.Director, os.ModePerm)
		if err != nil {
			return nil
		}
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.SG_BLOG_COFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
