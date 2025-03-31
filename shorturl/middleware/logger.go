package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		
		latency := time.Since(start)
		logrus.WithFields(logrus.Fields{
			"status":  c.Writer.Status(),
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
			"ip":      c.ClientIP(),
			"latency": latency,
		}).Info("请求处理完成")
	}
}