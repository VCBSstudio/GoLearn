package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimiter(rdb *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		key := "rate_limit:" + clientIP

		ctx := context.Background()
		
		// 使用Redis计数器
		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "内部服务错误"})
			return
		}

		// 如果是第一次请求，设置过期时间
		if count == 1 {
			rdb.Expire(ctx, key, window)
		}

		// 检查是否超过限制
		if count > int64(limit) {
			ttl, _ := rdb.TTL(ctx, key).Result()
			c.Header("X-RateLimit-Reset", strconv.FormatInt(int64(ttl/time.Second), 10))
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "请求过于频繁",
				"message": "请稍后再试",
			})
			return
		}

		c.Next()
	}
}