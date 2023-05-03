package models

type User struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null"`
	Email       string `json:"email" gorm:"uniqueIndex;not null"`
	Password    string `json:"password" gorm:"not null"`
	Address     string `json:"address" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
	ImageUrl    string `json:"image_url"`
}
