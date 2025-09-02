package models
import "gorm.io/gorm"
type Article struct {
    gorm.Model
    Title    string
    Category string // Jenis Kulit, Permasalahan, Faktor Eksternal
    Content  string `gorm:"type:text"`
    Author   string
}