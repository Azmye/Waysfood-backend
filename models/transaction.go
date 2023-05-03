package models

import "time"

type Transaction struct {
	ID              uint          `gorm:"primaryKey"`
	OrderID         uint          `gorm:"not null"`
	TransactionID   string        `gorm:"not null"`
	TransactionDate time.Time     `gorm:"not null"`
	Amount          float64       `gorm:"not null"`
	PaymentMethod   PaymentMethod `gorm:"not null"`
	Order           Order         `gorm:"foreignKey:OrderID"`
}
