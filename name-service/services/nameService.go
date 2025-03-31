package services

import (
	"context"
	"encoding/json"
	"fmt"
	"name-service/models"
	"name-service/utils"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type NameService struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewNameService(db *gorm.DB, redis *redis.Client) *NameService {
	return &NameService{
		db:    db,
		redis: redis,
	}
}

func (s *NameService) GenerateNames(info *models.BirthInfo) ([]models.NameSuggestion, error) {
	// 尝试从缓存获取
	cacheKey := fmt.Sprintf("names:%d:%d:%d:%d:%s", 
		info.Year, info.Month, info.Day, info.Hour, info.Gender)
	
	if cached, err := s.getFromCache(cacheKey); err == nil {
		return cached, nil
	}

	// 计算五行
	fiveElement := utils.CalculateFiveElements(
		info.Year,
		info.Month,
		info.Day,
		info.Hour,
	)

	// 获取适合的字符
	var characters []models.Character
	query := s.db.Where("gender IN (?)", []string{info.Gender, "B"}).
		Where("five_element = ?", fiveElement).
		Order("score DESC").
		Limit(100)
	
	if err := query.Find(&characters).Error; err != nil {
		return nil, err
	}

	// 生成名字建议
	suggestions := s.generateSuggestions(characters, info)

	// 保存到缓存
	s.saveToCache(cacheKey, suggestions)

	return suggestions, nil
}

func (s *NameService) generateSuggestions(chars []models.Character, info *models.BirthInfo) []models.NameSuggestion {
	suggestions := make([]models.NameSuggestion, 0)
	
	// 根据五行相生相克规则筛选字符
	compatibleChars := utils.FilterCompatibleCharacters(chars, info)
	
	// 生成双字名
	for i := 0; i < len(compatibleChars); i++ {
		for j := i + 1; j < len(compatibleChars); j++ {
			score := s.calculateNameScore(
				compatibleChars[i],
				compatibleChars[j],
				info,
			)
			
			if score >= 70 { // 只推荐评分较高的名字
				suggestion := models.NameSuggestion{
					FirstName:   "张", // 这里应该是用户输入的姓氏
					LastName:    compatibleChars[i].Character + compatibleChars[j].Character,
					Meaning:     fmt.Sprintf("%s，%s", compatibleChars[i].Meaning, compatibleChars[j].Meaning),
					FiveElement: compatibleChars[i].FiveElement + "," + compatibleChars[j].FiveElement,
					Score:       score,
				}
				suggestions = append(suggestions, suggestion)
			}
		}
	}

	return suggestions
}

func (s *NameService) calculateNameScore(char1, char2 models.Character, info *models.BirthInfo) int {
	// 基础分数
	baseScore := (char1.Score + char2.Score) / 2
	
	// 计算八字评分
	baziScore := utils.CalculateBaZiScore(info, char1.FiveElement, char2.FiveElement)
	
	// 计算笔画评分
	strokesScore := utils.CalculateStrokesScore(char1.Strokes + char2.Strokes)
	
	// 综合评分：基础分 40%，八字评分 40%，笔画评分 20%
	finalScore := int(float64(baseScore)*0.4 + float64(baziScore)*0.4 + float64(strokesScore)*0.2)
	
	return finalScore
}

func (s *NameService) getFromCache(key string) ([]models.NameSuggestion, error) {
	ctx := context.Background()
	data, err := s.redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var suggestions []models.NameSuggestion
	err = json.Unmarshal(data, &suggestions)
	return suggestions, err
}

func (s *NameService) saveToCache(key string, suggestions []models.NameSuggestion) error {
	ctx := context.Background()
	data, err := json.Marshal(suggestions)
	if err != nil {
		return err
	}

	return s.redis.Set(ctx, key, data, 24*time.Hour).Err()
}