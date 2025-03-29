package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", Login)
		userGroup.POST("/register", Register)
		userGroup.GET("/profile", Profile)
	}
}
