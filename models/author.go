package models

type Author struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}
