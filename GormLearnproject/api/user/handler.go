package user

import (
    "GormLearnproject/model"
    "github.com/gin-gonic/gin"
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