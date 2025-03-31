package utils

var (
	heavenlyStems   = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	earthlyBranches = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	fiveElements    = []string{"金", "木", "水", "火", "土"}
)

func CalculateFiveElements(year, month, day, hour int) string {
	// 简化版五行计算
	baseYear := 1900
	stemIndex := (year - baseYear) % 10
	branchIndex := (year - baseYear) % 12

	// 根据时辰调整五行
	hourIndex := hour / 2

	// 简单的五行计算规则
	elementIndex := (stemIndex + branchIndex + hourIndex) % 5
	return fiveElements[elementIndex]
}
