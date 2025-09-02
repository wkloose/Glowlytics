package models
import "gorm.io/gorm"
type Lifestyle struct {
    gorm.Model
    UserID   uint
    User     User
    Sleep    int    // jam tidur rata-rata
    Diet     string // contoh: "high sugar", "balanced"
    Exercise string // contoh: "3x a week"
    Stress   string // contoh: "low/medium/high"
    SunExposure string
}