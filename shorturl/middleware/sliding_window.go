package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SlidingWindowLimiter(rdb *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now().UnixNano()
		clientIP := c.ClientIP()
		key := fmt.Sprintf("sliding:%s", clientIP)

		// 移除时间窗口外的请求
		min := now - window.Nanoseconds()

		// 使用Redis事务
		pipe := rdb.TxPipeline()
		pipe.ZRemRangeByScore(c.Request.Context(), key, "0", fmt.Sprintf("%d", min))
		pipe.ZAdd(c.Request.Context(), key, redis.Z{
			Score:  float64(now),
			Member: now,
		})
		pipe.Expire(c.Request.Context(), key, window)
		count := pipe.ZCard(c.Request.Context(), key)
		if _, err := pipe.Exec(c.Request.Context()); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "内部服务错误"})
			return
		}

		if count.Val() > int64(limit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "请求过于频繁",
				"message": "请稍后再试",
			})
			return
		}

		c.Next()
	}
}
