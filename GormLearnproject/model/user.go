package model

import (
    "gorm.io/gorm"
    "time"
)

type User struct {
    ID        uint           `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
    Name      string         `gorm:"type:varchar(20);not null"`
    Email     string         `gorm:"type:varchar(100);uniqueIndex"`
    Password  string         `gorm:"type:varchar(100);not null"`
    Age       int            `gorm:"type:int;not null"`
}