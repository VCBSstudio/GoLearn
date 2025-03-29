package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name     string `gorm:"type:varchar(20);not null"`
    Age      int    `gorm:"type:int;not null"`
    Email    string `gorm:"type:varchar(100);uniqueIndex"`
}