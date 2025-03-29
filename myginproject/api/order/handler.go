package order

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateOrder(c *gin.Context) {
	var req CreateOrderRequest

	if err := c.BindJSON(&req); err != nil {
		// 类型断言判断是否为验证错误
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make([]string, 0)
			for _, e := range validationErrors {
				switch e.Tag() {
				case "required":
					errorMessages = append(errorMessages, e.Field()+" 是必填字段")
				case "gt":
					errorMessages = append(errorMessages, e.Field()+" 必须大于 0")
				}
			}
			c.JSON(400, gin.H{
				"errors": errorMessages,
			})
			return
		}

		// 其他类型错误
		c.JSON(400, gin.H{
			"error": "无效的请求数据",
		})
		return
	}

	// 处理业务逻辑...
	c.JSON(200, gin.H{
		"message": "订单创建成功",
		"order":   req,
	})
}
