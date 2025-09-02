package models

import "gorm.io/gorm"
type ScanHistory struct {
    gorm.Model
    UserID      uint
    User        User
    ImageURL    string `gorm:"not null"`
    Diagnosis   string // contoh: "Acne Vulgaris"
    Severity    string // Mild / Moderate / Severe
    Suggestion  string
}