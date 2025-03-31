package utils

import (
	"name-service/models"
	"time"
)

// 天干地支对应的五行
var (
	stemElements = map[string]string{
		"甲": "木", "乙": "木",
		"丙": "火", "丁": "火",
		"戊": "土", "己": "土",
		"庚": "金", "辛": "金",
		"壬": "水", "癸": "水",
	}

	branchElements = map[string]string{
		"子": "水", "丑": "土", "寅": "木",
		"卯": "木", "辰": "土", "巳": "火",
		"午": "火", "未": "土", "申": "金",
		"酉": "金", "戌": "土", "亥": "水",
	}

	// 五行相生关系权重
	elementGenerationWeight = 1.2
	// 五行相克关系权重
	elementControlWeight = 0.8
)

// 计算八字评分
func CalculateBaZiScore(info *models.BirthInfo, element1, element2 string) int {
	year, month, day := info.Year, info.Month, info.Day

	// 获取年柱
	yearStem, yearBranch := calculateYearPillar(year)
	// 获取月柱
	monthStem, monthBranch := calculateMonthPillar(year, month)
	// 获取日柱
	dayStem, dayBranch := calculateDayPillar(year, month, day)
	// 获取时柱
	hourStem, hourBranch := calculateHourPillar(day, info.Hour)

	// 计算八字五行分布
	elements := map[string]int{
		"金": 0, "木": 0, "水": 0, "火": 0, "土": 0,
	}

	// 统计八字中各五行的数量
	elements[stemElements[yearStem]]++
	elements[stemElements[monthStem]]++
	elements[stemElements[dayStem]]++
	elements[stemElements[hourStem]]++
	elements[branchElements[yearBranch]]++
	elements[branchElements[monthBranch]]++
	elements[branchElements[dayBranch]]++
	elements[branchElements[hourBranch]]++

	// 计算五行缺失和过旺
	score := calculateElementBalance(elements)

	// 考虑名字五行与八字五行的配合
	score += calculateElementCompatibility(elements, element1, element2)

	return normalizeScore(score)
}

// 计算五行平衡度
func calculateElementBalance(elements map[string]int) int {
	score := 100
	avg := float64(8) / 5 // 理想情况下每个五行平均分布

	for _, count := range elements {
		diff := float64(count) - avg
		if diff > 0 {
			// 五行过旺，扣分
			score -= int(diff * 10)
		} else if diff < 0 {
			// 五行不足，扣分
			score -= int(-diff * 10)
		}
	}
	return score
}

// 计算名字五行与八字五行的配合度
func calculateElementCompatibility(baziElements map[string]int, element1, element2 string) int {
	score := 0

	// 补充八字中缺失的五行
	if baziElements[element1] == 0 {
		score += 20
	}
	if baziElements[element2] == 0 {
		score += 20
	}

	// 考虑五行相生
	if IsElementsCompatible(element1, element2) {
		score += 15
	}

	return score
}

// 标准化分数到0-100范围
func normalizeScore(score int) int {
	if score < 0 {
		return 0
	}
	if score > 100 {
		return 100
	}
	return score
}

// 计算年柱
func calculateYearPillar(year int) (string, string) {
	// 1900年为庚子年
	baseYear := 1900
	stemIndex := (year - baseYear) % 10
	branchIndex := (year - baseYear) % 12

	if stemIndex < 0 {
		stemIndex += 10
	}
	if branchIndex < 0 {
		branchIndex += 12
	}

	return heavenlyStems[stemIndex], earthlyBranches[branchIndex]
}

// 计算月柱
func calculateMonthPillar(year, month int) (string, string) {
	// 确定月干的起始偏移
	yearStem, _ := calculateYearPillar(year)
	stemOffset := 0
	for i, stem := range heavenlyStems {
		if stem == yearStem {
			stemOffset = i
			break
		}
	}

	// 计算月干和月支
	monthStemIndex := (stemOffset*2 + month + 1) % 10
	monthBranchIndex := (month + 1) % 12

	return heavenlyStems[monthStemIndex], earthlyBranches[monthBranchIndex]
}

// 计算日柱
func calculateDayPillar(year, month, day int) (string, string) {
	// 使用基姆拉尔森计算公式计算日柱
	// 1900年1月1日为甲戌日
	baseDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	currentDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	days := int(currentDate.Sub(baseDate).Hours() / 24)
	stemIndex := (days + 10) % 10
	branchIndex := (days + 10) % 12

	return heavenlyStems[stemIndex], earthlyBranches[branchIndex]
}

// 计算时柱
func calculateHourPillar(day int, hour int) (string, string) {
	// 子时为晚上23点到凌晨1点
	hourBranchIndex := ((hour + 1) / 2) % 12
	hourStemIndex := (day*2 + hourBranchIndex) % 10

	return heavenlyStems[hourStemIndex], earthlyBranches[hourBranchIndex]
}
