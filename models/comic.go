package models

import (
	"gorm.io/gorm"
)

type Comic struct {
	gorm.Model
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	AuthorID    uint     `json:"author_id"`
	Author      Author   `gorm:"foreignKey:AuthorID" json:"author"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category"`
}
	