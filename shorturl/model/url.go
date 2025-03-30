package model

import (
    "gorm.io/gorm"
    "time"
)

type URL struct {
    gorm.Model
    OriginalURL string    `gorm:"type:text;not null"`
    ShortCode   string    `gorm:"type:varchar(10);uniqueIndex;not null"`
    Visits      int64     `gorm:"default:0"`
    LastVisit   time.Time `gorm:"default:NULL"`  // 修改这里
    UserID      uint      `gorm:"index"`
}