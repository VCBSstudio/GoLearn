package config

import (
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	// 从环境变量获取Redis配置
	host := getEnv("REDIS_HOST", "localhost")
	port := getEnv("REDIS_PORT", "6379")
	password := getEnv("REDIS_PASSWORD", "your_strong_password")

	return redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
}
