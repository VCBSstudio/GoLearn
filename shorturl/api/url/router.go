package url

import (
    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"
    "gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, rdb *redis.Client) {
    handler := NewHandler(db, rdb)

    r.POST("/api/urls", handler.Create)
    r.GET("/api/urls/:code/stats", handler.Stats)
    r.GET("/:code", handler.Redirect)
}