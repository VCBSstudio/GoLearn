package utils

import (
	"name-service/models"
)

func FilterCompatibleCharacters(chars []models.Character, info *models.BirthInfo) []models.Character {
	var result []models.Character
	birthElement := CalculateFiveElements(info.Year, info.Month, info.Day, info.Hour)

	for _, char := range chars {
		if IsElementsCompatible(birthElement, char.FiveElement) {
			result = append(result, char)
		}
	}

	return result
}

func IsElementsCompatible(element1, element2 string) bool {
	// 五行相生: 金生水，水生木，木生火，火生土，土生金
	generationMap := map[string]string{
		"金": "水",
		"水": "木",
		"木": "火",
		"火": "土",
		"土": "金",
	}

	return generationMap[element1] == element2 || generationMap[element2] == element1
}

func CalculateStrokesScore(totalStrokes int) int {
	// 根据笔画数计算分数，这里是简化版
	goodStrokes := []int{1, 3, 5, 7, 8, 11, 13, 15, 16, 21, 23, 24, 29, 31, 32, 33, 35, 37, 39, 41, 45, 47}

	for _, stroke := range goodStrokes {
		if totalStrokes == stroke {
			return 90
		}
	}

	return 70
}

// func CalculateBaZiScore(info *models.BirthInfo, element1, element2 string) int {
// 	// 这里可以实现更复杂的八字评分算法
// 	return 80
// }
