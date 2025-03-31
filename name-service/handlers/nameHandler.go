package handlers

import (
	"name-service/models"
	"name-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NameHandler struct {
	nameService *services.NameService
}

func NewNameHandler(nameService *services.NameService) *NameHandler {
	return &NameHandler{
		nameService: nameService,
	}
}

func (h *NameHandler) GenerateNames(c *gin.Context) {
	var birthInfo models.BirthInfo
	if err := c.ShouldBindJSON(&birthInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数无效",
		})
		return
	}

	suggestions, err := h.nameService.GenerateNames(&birthInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "生成名字失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"names": suggestions,
	})
}
