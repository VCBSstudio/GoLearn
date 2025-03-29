package user

import (
	"GormLearnproject/middleware"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	handler := NewHandler(db, rdb)

	// 公开路由
	r.POST("/login", handler.Login)
	r.PATCH("/users/:id/password", handler.UpdatePassword)

	// 需要认证的路由
	userGroup := r.Group("/users")
	userGroup.Use(middleware.JWTAuth())
	{
		userGroup.POST("", handler.Create)
		userGroup.GET("", handler.List)
	}
}
