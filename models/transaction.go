package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionID int    `gorm:"not null"`
	ProductID     uint   `gorm:"not null"`
	CustomerID    uint   `gorm:"not null"`
	Status        string `gorm:"not null"`
	Location      string `gorm:"not null"`
	TotalPrice    int    `gorm:"not null"`
	Product       Product
	Customer      User
}
