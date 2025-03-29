package user

import (
    "GormLearnproject/model"
    "GormLearnproject/utils"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type Handler struct {
    db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
    return &Handler{db: db}
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
    var users []model.User
    result := h.db.Find(&users)
    if result.Error != nil {
        c.JSON(500, gin.H{"error": result.Error.Error()})
        return
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
            "id": user.ID,
            "name": user.Name,
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