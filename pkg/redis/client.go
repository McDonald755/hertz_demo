package redis

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	redisv8 "github.com/go-redis/redis/v8"
	"hertz_demo/pkg/viper"
)

var (
	Client *redisv8.Client
	config *viper.Redis
)

func InitRedis() {
	config = viper.Conf.Redis
	initRedis(context.Background())
}

func initRedis(ctx context.Context) {
	rdb := redisv8.NewClient(&redisv8.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
	})
	if rdb == nil {
		hlog.CtxFatalf(ctx, "[Redis] Init Failed")
	}
	hlog.CtxInfof(ctx, "[Redis] PING: %s\n", Client.Ping(ctx))

	//暴露redis连接
	Client = rdb
}

// GetValue 获取redis value
func GetValue(ctx context.Context, key string) string {
	res, err := Client.Get(ctx, key).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[Redis] Get redis value failed, key: %v, err: %v", key, err)
		return ""
	}
	return res
}
