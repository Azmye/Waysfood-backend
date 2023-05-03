package models

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" form:"name" gorm:"not null"`
	Description string  `json:"description" form:"description" gorm:"not null"`
	Price       float64 `json:"price" form:"price" gorm:"not null"`
	ImageURL    string  `json:"image_url" form:"image_url" gorm:"not null"`
}
