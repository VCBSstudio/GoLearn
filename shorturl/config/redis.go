package config

import (
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "your_strong_password", // Redis 密码
		DB:       0,
	})
}
