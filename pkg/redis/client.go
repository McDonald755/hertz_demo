package redis

import (
	"context"
	"demo2/pkg/viper"
	redisv8 "github.com/go-redis/redis/v8"
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

}
