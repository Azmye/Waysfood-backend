package models

import "time"

type Order struct {
	ID             uint           `gorm:"primaryKey"`
	UserID         uint           `gorm:"not null"`
	OrderDate      time.Time      `gorm:"not null"`
	Total          float64        `gorm:"not null"`
	PaymentStatus  PaymentStatus  `gorm:"not null;default:'pending'"`
	PaymentMethod  PaymentMethod  `gorm:"not null"`
	User           User           `gorm:"foreignKey:UserID"`
	Details        []OrderDetails `gorm:"foreignKey:OrderID"`
	Transactions   []Transaction  `gorm:"foreignKey:OrderID"`
	MidtransCharge MidtransCharge `gorm:"embedded;embeddedPrefix:midtrans_"`
}
