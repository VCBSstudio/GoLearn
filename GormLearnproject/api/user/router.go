package user

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    handler := NewHandler(db)
    
    userGroup := r.Group("/users")
    {
        userGroup.POST("", handler.Create)
        userGroup.GET("", handler.List)
    }
}