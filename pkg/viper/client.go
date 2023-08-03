package viper

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
	"hertz_demo/biz/utils"
	"os"
	"path/filepath"
)

var Conf *Config //定义全局变量

func InitViper() {
	ctx := context.Background()
	viper.SetConfigType("yaml") //设置类型
	runEnv := os.Getenv("RUN_ENV")
	confPath := utils.GetConfAbPath() //获取配置文件路径路径

	//判断运行环境
	if runEnv == "DEV" {
		viper.SetConfigFile(filepath.Join(confPath, "dev.config.yaml"))
	} else if runEnv == "PROD" {
		viper.SetConfigFile(filepath.Join(confPath, "prod.config.yaml"))
	} else {
		viper.SetConfigFile(filepath.Join(confPath, "default.config.yaml"))
	}

	if err := viper.ReadInConfig(); err != nil {
		hlog.CtxErrorf(ctx, "[Viper] ReadInConfig failed, err: %v", err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		hlog.CtxErrorf(ctx, "[Viper] Unmarshal failed, err: %v", err)
	}

	hlog.CtxInfof(ctx, "[Viper] Conf.App: %#v", Conf.App)
	hlog.CtxInfof(ctx, "[Viper] Conf.Redis: %#v", Conf.Redis)
	hlog.CtxInfof(ctx, "[Viper] Conf.Mysql: %#v", Conf.Mysql)

}
