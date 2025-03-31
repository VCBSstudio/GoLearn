package url

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"shorturl/model"
	"shorturl/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Handler struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewHandler(db *gorm.DB, rdb *redis.Client) *Handler {
	return &Handler{db: db, rdb: rdb}
}

func (h *Handler) Create(c *gin.Context) {
	var input struct {
		URL string `json:"url" binding:"required,url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, 400, "无效的URL格式")
		return
	}

	hash := md5.Sum([]byte(input.URL + time.Now().String()))
	shortCode := base64.URLEncoding.EncodeToString(hash[:])[:6]

	url := model.URL{
		OriginalURL: input.URL,
		ShortCode:   shortCode,
	}

	if err := h.db.Create(&url).Error; err != nil {
		utils.Logger.WithError(err).Error("创建短链接失败 - 数据库错误")
		print("lt -- 为什么创建失败 ")
		utils.ErrorResponse(c, 500, "创建短链接失败")
		return
	}

	ctx := context.Background()
	h.rdb.Set(ctx, "url:"+shortCode, input.URL, 24*time.Hour)

	c.JSON(200, gin.H{
		"short_url": fmt.Sprintf("http://%s/%s", c.Request.Host, shortCode),
		"code":      shortCode,
	})
}

func (h *Handler) Redirect(c *gin.Context) {
	// 使用Redis缓存热门URL
	code := c.Param("code")
	ctx := context.Background()

	// 先检查Redis缓存
	originalURL, err := h.rdb.Get(ctx, "url:"+code).Result()
	if err == nil {
		go h.updateVisits(code)
		c.Redirect(302, originalURL)
		return
	}

	// 数据库查询使用Select只获取必要字段
	var url model.URL
	if err := h.db.Select("original_url").Where("short_code = ?", code).First(&url).Error; err != nil {
		utils.ErrorResponse(c, 404, "短链接不存在")
		return
	}

	// 缓存到Redis
	h.rdb.Set(ctx, "url:"+code, url.OriginalURL, 24*time.Hour)
	go h.updateVisits(code)
	c.Redirect(302, url.OriginalURL)
}

func (h *Handler) Stats(c *gin.Context) {
	code := c.Param("code")

	var url model.URL
	if err := h.db.Where("short_code = ?", code).First(&url).Error; err != nil {
		utils.ErrorResponse(c, 404, "短链接不存在")
		return
	}

	c.JSON(200, gin.H{
		"original_url": url.OriginalURL,
		"visits":       url.Visits,
		"last_visit":   url.LastVisit,
		"created_at":   url.CreatedAt,
	})
}

func (h *Handler) updateVisits(code string) {
	h.db.Model(&model.URL{}).
		Where("short_code = ?", code).
		Updates(map[string]interface{}{
			"visits":     gorm.Expr("visits + 1"),
			"last_visit": gorm.Expr("CURRENT_TIMESTAMP"),  // 修改这里
		})
}
