package models

type PaymentStatus int

const (
	PaymentPending PaymentStatus = iota
	PaymentSuccess
	PaymentFailed
)

func (s PaymentStatus) String() string {
	return [...]string{"Pending", "Success", "Failed"}[s]
}
