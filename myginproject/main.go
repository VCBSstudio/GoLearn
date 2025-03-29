// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"myginproject/middleware"
// )

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	r.POST("/create", func(c *gin.Context) {
// 		var json map[string]interface{}
// 		if err := c.ShouldBindJSON(&json); err != nil {
// 			c.JSON(400, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(200, gin.H{"status": "ok", "data": json})
// 	})

// 	r.Run(":9000")

// }

package main

import (
	"myginproject/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // 包含了 Recovery 中间件

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 这个处理函数可能会发生 panic
	// r.GET("/dangerous", func(c *gin.Context) {
	// 	// 模拟一个 panic 情况
	// 	var slice []string
	// 	// 这里会引发 panic，因为 slice 是 nil
	// 	slice[0] = "hello" // 这会导致程序崩溃

	// 	c.JSON(200, gin.H{"message": "你永远看不到这条消息"})
	// })

	// 使用自定义的 Recovery 中间件
	r.Use(middleware.SimpleRecovery())

	r.GET("/dangerous", func(c *gin.Context) {
		var slice []string
		slice[0] = "hello" // 这里会 panic

		c.JSON(200, gin.H{"message": "正常响应"})
	})

	r.POST("/create", func(c *gin.Context) {
		var json map[string]interface{}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "data": json})
	})

	r.Run(":9000")
}
