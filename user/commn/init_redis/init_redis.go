package init_redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Init_redis() *redis.Client {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	if err := rdb.Ping(ctx).Err(); err != nil {

	}
	zap.L().Info("redis init success")
	return rdb
}
