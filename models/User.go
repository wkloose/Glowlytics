package models

import "gorm.io/gorm"
type User struct {
    gorm.Model
    Name     string `gorm:"not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Age      int
    Gender   string
    SkinType string // oily, dry, combination, sensitive
}