package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "your_strong_password", // Redis 密码
		DB:       0,
	})

	// 测试连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("Redis 连接失败: " + err.Error())
	}

	return rdb
}
