package models

type BirthInfo struct {
	Year   int    `json:"year"`
	Month  int    `json:"month"`
	Day    int    `json:"day"`
	Hour   int    `json:"hour"`
	Minute int    `json:"minute"`
	Gender string `json:"gender"`
}

type Name struct {
	ID          uint   `gorm:"primarykey"`
	Character   string `gorm:"type:varchar(10);not null"`
	Meaning     string `gorm:"type:text"`
	FiveElement string `gorm:"type:varchar(10)"` // 金木水火土
	Strokes     int    `gorm:"type:int"`
}

type NameSuggestion struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Meaning     string `json:"meaning"`
	FiveElement string `json:"fiveElement"`
	Score       int    `json:"score"`
}
