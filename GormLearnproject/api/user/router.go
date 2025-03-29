package user

import (
    "GormLearnproject/middleware"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    handler := NewHandler(db)
    
    // 公开路由 - 不需要 token
    r.POST("/login", handler.Login)
    r.PATCH("/users/:id/password", handler.UpdatePassword)
    
    // 需要认证的路由 - 必须携带有效的 token
    userGroup := r.Group("/users")
    userGroup.Use(middleware.JWTAuth())  // 这里添加了 JWT 认证中间件
    {
        userGroup.POST("", handler.Create)
        userGroup.GET("", handler.List)
    }
}