package cartDto

type CreateCartRequest struct {
	ProductID  uint   `json:"product_id" form:"product_id" gorm:"not null"`
	CustomerID uint   `json:"customer_id" form:"customer_id" gorm:"not null"`
	Quantity   uint   `json:"qty" form:"qty" gorm:"not null"`
	Status     string `json:"status" form:"qty" gorm:"not null"`
}

type UpdateCartRequest struct {
	ProductID  uint   `json:"product_id"`
	CustomerID uint   `json:"customer_id"`
	Quantity   uint   `json:"qty"`
	Status     string `json:"status"`
}
