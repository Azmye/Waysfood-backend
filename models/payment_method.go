package models

import "gorm.io/gorm"

type PaymentMethod int

const (
	PaymentMethodCreditCard PaymentMethod = iota
	PaymentMethodBankTransfer
	PaymentMethodVirtualAccount
)

func (p PaymentMethod) String() string {
	return [...]string{"CreditCard", "BankTransfer", "VirtualAccount"}[p]
}

type MidtransCharge struct {
	gorm.Model
	OrderID       uint
	TransactionID string
	Status        string
	PaymentMethod PaymentMethod
	PaymentType   string
	VaNumber      string
	Bank          string
	GrossAmount   float64
}
