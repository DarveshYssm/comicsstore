package models

type Comic struct {
	ID          uint     `json:"id" gorm:"primary_key"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	AuthorID    uint     `json:"author_id"`
	CategoryID  uint     `json:"category_id"`
	Author      Author   `json:"author" gorm:"foreignkey:AuthorID"`
	Category    Category `json:"category" gorm:"foreignkey:CategoryID"`
}
