package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func TokenBucketLimiter(rdb *redis.Client, rate int, capacity int) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "token_bucket:" + c.ClientIP()
		now := time.Now().Unix()

		// 使用Lua脚本保证原子性
		lua := `
		local key = KEYS[1]
		local now = tonumber(ARGV[1])
		local rate = tonumber(ARGV[2])
		local capacity = tonumber(ARGV[3])
		
		local last = redis.call("GET", key)
		local tokens = capacity
		
		if last then
			local elapsed = now - tonumber(last)
			local new_tokens = math.floor(elapsed * rate)
			tokens = math.min(capacity, new_tokens + (redis.call("GET", key..":tokens") or capacity))
		end
		
		if tokens <= 0 then
			return 0
		end
		
		redis.call("SET", key, now)
		redis.call("SET", key..":tokens", tokens - 1)
		return 1
		`

		val, err := rdb.Eval(c.Request.Context(), lua, []string{key}, now, rate, capacity).Result()
		if err != nil || val.(int64) == 0 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "请求过于频繁",
				"message": "请稍后再试",
			})
			return
		}

		c.Next()
	}
}
