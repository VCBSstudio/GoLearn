package user

import (
	"GormLearnproject/model"
	"GormLearnproject/utils"
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
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

// 删除重复的 import 声明
// 直接开始错误响应函数的定义
func errorResponse(c *gin.Context, code int, message string) {
	utils.Logger.WithFields(logrus.Fields{
		"status": code,
		"error":  message,
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	}).Error("API错误")

	c.JSON(code, gin.H{"error": message})
}

// 修改 List 函数
func (h *Handler) List(c *gin.Context) {
	ctx := context.Background()
	cacheKey := "users:all:v1"

	utils.Logger.WithFields(logrus.Fields{
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	}).Info("获取用户列表")

	// 尝试从缓存获取
	cachedData, err := h.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var users []model.User
		if err := json.Unmarshal([]byte(cachedData), &users); err == nil {
			utils.Logger.Info("从Redis缓存获取数据")
			c.JSON(200, users)
			return
		}
	}

	// 从数据库查询
	utils.Logger.Info("从数据库查询数据")
	var users []model.User
	result := h.db.Find(&users)
	if result.Error != nil {
		errorResponse(c, 500, "数据库查询失败")
		return
	}

	// 更新缓存
	if usersJson, err := json.Marshal(users); err == nil {
		h.rdb.Set(ctx, cacheKey, usersJson, 30*time.Minute)
		utils.Logger.Info("数据已更新到Redis缓存")
	}

	c.JSON(200, users)
}

// 修改 Login 函数
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// 保留这个新版本的 Login 函数，删除旧版本
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, 400, "无效的请求参数")
		return
	}

	utils.Logger.WithFields(logrus.Fields{
		"email": req.Email,
	}).Info("用户登录尝试")

	var user model.User
	if err := h.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		errorResponse(c, 401, "用户名或密码错误")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.Logger.WithFields(logrus.Fields{
			"email": req.Email,
		}).Warn("密码验证失败")
		errorResponse(c, 401, "用户名或密码错误")
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		errorResponse(c, 500, "生成令牌失败")
		return
	}

	utils.Logger.WithFields(logrus.Fields{
		"userId": user.ID,
		"email":  user.Email,
	}).Info("用户登录成功")

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
