package user

import "github.com/gin-gonic/gin"

// 定义请求结构体
type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
    var req LoginRequest
    
    // 绑定 JSON 数据
    if err := c.BindJSON(&req); err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
    
    // 使用绑定后的数据
    c.JSON(200, gin.H{
        "message": "登录成功",
        "username": req.Username,
    })
}

func Register(c *gin.Context) {
	c.JSON(200, gin.H{"message": "register success"})
}

func Profile(c *gin.Context) {
	c.JSON(200, gin.H{"message": "user profile"})
}
