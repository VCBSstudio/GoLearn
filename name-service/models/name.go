package models

type BirthInfo struct {
	Year   int    `json:"year" binding:"required"`
	Month  int    `json:"month" binding:"required,min=1,max=12"`
	Day    int    `json:"day" binding:"required,min=1,max=31"`
	Hour   int    `json:"hour" binding:"required,min=0,max=23"`
	Minute int    `json:"minute" binding:"required,min=0,max=59"`
	Gender string `json:"gender" binding:"required,oneof=M F"`
}

type Name struct {
	ID          uint   `gorm:"primarykey"`
	Character   string `gorm:"type:varchar(10);not null"`
	Meaning     string `gorm:"type:text"`
	FiveElement string `gorm:"type:varchar(10)"` // 金木水火土
	Strokes     int    `gorm:"type:int"`
}

type NameSuggestion struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Meaning    string `json:"meaning"`
	FiveElement string `json:"fiveElement"`
	Score      int    `json:"score"`
}