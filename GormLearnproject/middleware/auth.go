package middleware

import (
    "GormLearnproject/utils"
    "github.com/gin-gonic/gin"
    "strings"
)

func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "未提供认证令牌"})
            c.Abort()
            return
        }

        // 从 Bearer Token 中提取令牌
        parts := strings.SplitN(authHeader, " ", 2)
        if !(len(parts) == 2 && parts[0] == "Bearer") {
            c.JSON(401, gin.H{"error": "认证格式错误"})
            c.Abort()
            return
        }

        claims, err := utils.ParseToken(parts[1])
        if err != nil {
            c.JSON(401, gin.H{"error": "无效的令牌"})
            c.Abort()
            return
        }

        // 将用户ID存储在上下文中
        c.Set("userID", claims.UserID)
        c.Next()
    }
}