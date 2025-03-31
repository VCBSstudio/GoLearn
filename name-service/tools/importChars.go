package tools

import (
	"encoding/json"
	"name-service/models"
	"os"

	"gorm.io/gorm"
)

type CharacterData struct {
	Characters []struct {
		Char        string   `json:"char"`
		Strokes     int      `json:"strokes"`
		FiveElement string   `json:"five_element"`
		Meaning     string   `json:"meaning"`
		Score       int      `json:"score"`
		Gender      string   `json:"gender"`
		Luck        []string `json:"luck"`
	} `json:"characters"`
}

func ImportCharacters(db *gorm.DB, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var charData CharacterData
	if err := json.Unmarshal(data, &charData); err != nil {
		return err
	}

	for _, c := range charData.Characters {
		char := models.Character{
			Character:   c.Char,
			Strokes:     c.Strokes,
			FiveElement: c.FiveElement,
			Meaning:     c.Meaning,
			Score:       c.Score,
			Gender:      c.Gender,
		}
		if err := db.Create(&char).Error; err != nil {
			return err
		}
	}
	return nil
}
