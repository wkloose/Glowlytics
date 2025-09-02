package models
import "gorm.io/gorm"
type Recommendation struct {
    gorm.Model
    UserID   uint
    User     User
    Products string // JSON: cleanser, sunscreen, moisturizer
    Routine  string // JSON/string detail rutinitas harian
    Advice   string // lifestyle tips
}