package user

import (
	"GormLearnproject/model"
	"GormLearnproject/utils"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	db  *gorm.DB
	rdb *redis.Client // 添加 redis 客户端
}

func NewHandler(db *gorm.DB, rdb *redis.Client) *Handler {
	return &Handler{
		db:  db,
		rdb: rdb,
	}
}

func (h *Handler) Create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := h.db.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, user)
}

func (h *Handler) List(c *gin.Context) {
	ctx := context.Background()
	cacheKey := "users:all"

	// 尝试从缓存获取
	cachedData, err := h.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		// 缓存命中
		var users []model.User
		if err := json.Unmarshal([]byte(cachedData), &users); err == nil {
			fmt.Println("从 Redis 缓存中获取数据") // 添加日志
			c.JSON(200, users)
			return
		}
	}

	// 缓存未命中，从数据库查询
	fmt.Println("从数据库中获取数据") // 添加日志
	var users []model.User
	result := h.db.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// 将数据存入缓存
	if usersJson, err := json.Marshal(users); err == nil {
		h.rdb.Set(ctx, cacheKey, usersJson, 5*time.Minute)
		fmt.Println("数据已存入 Redis 缓存") // 添加日志
	}

	c.JSON(200, users)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := h.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "用户名或密码错误"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(401, gin.H{"error": "用户名或密码错误"})
		return
	}

	// 生成Token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "生成令牌失败"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

type UpdatePasswordRequest struct {
	Password string `json:"password" binding:"required,min=6"`
}

func (h *Handler) UpdatePassword(c *gin.Context) {
	var req UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.Param("id")

	// 生成密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "密码加密失败"})
		return
	}

	// 更新用户密码
	result := h.db.Model(&model.User{}).Where("id = ?", userID).Update("password", string(hashedPassword))
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "更新密码失败"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "用户不存在"})
		return
	}

	c.JSON(200, gin.H{"message": "密码更新成功"})
}
