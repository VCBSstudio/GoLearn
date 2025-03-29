package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SimpleRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		// defer 确保这个函数在当前函数返回前一定会执行
		defer func() {
			// recover() 可以捕获到 panic 传递的错误信息
			if err := recover(); err != nil {
				// 打印错误信息
				fmt.Printf("发生了 panic: %v\n", err)

				// 返回 500 错误给客户端
				c.AbortWithStatusJSON(500, gin.H{
					"error": "服务器内部错误",
				})
			}
		}()

		// 继续执行后续的处理函数
		c.Next()
	}
}
