package models

type Character struct {
	ID          uint   `gorm:"primarykey"`
	Character   string `gorm:"type:varchar(1);unique;not null"`
	Strokes     int    `gorm:"type:int;not null"`
	FiveElement string `gorm:"type:varchar(10);not null"`
	Meaning     string `gorm:"type:text"`
	Score       int    `gorm:"type:int;not null"`
	Gender      string `gorm:"type:varchar(1)"` // M, F, B(both)
	Usage       int    `gorm:"type:int;default:0"` // 使用频率
}